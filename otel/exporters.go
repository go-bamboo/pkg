package otel

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/net/ip"
	prom "github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/contrib/propagators/aws/xray"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Provider struct {
	tracer   *tracesdk.TracerProvider
	meter    *metric.MeterProvider
	traceFn  *os.File
	metricFn *os.File
}

func (o *Provider) Close() {
	if o.tracer != nil {
		o.tracer.Shutdown(context.TODO())
	}
	if o.meter != nil {
		o.meter.Shutdown(context.TODO())
	}
	if o.traceFn != nil {
		o.traceFn.Sync()
		o.traceFn.Close()
	}
	if o.metricFn != nil {
		o.metricFn.Sync()
		o.metricFn.Close()
	}
}

func MustNewProvider(c *Conf, serviceName string, uuid string) *Provider {
	tp, err := NewProvider(c, serviceName, uuid)
	if err != nil {
		log.Fatal(err)
	}
	return tp
}

// NewProvider Get trace provider
func NewProvider(c *Conf, serviceName string, uuid string) (*Provider, error) {
	tp, err := newTracerProvider(c, serviceName, uuid)
	if err != nil {
		return nil, err
	}
	sdk, err := newMeterProvider(c, serviceName, uuid)
	if err != nil {
		return nil, err
	}
	otel.SetTracerProvider(tp)
	otel.SetMeterProvider(sdk)
	return &Provider{
		tracer: tp,
		meter:  sdk,
	}, nil
}

type noOutput int

func (*noOutput) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func newTracerProvider(c *Conf, serviceName string, uuid string) (tp *tracesdk.TracerProvider, err error) {
	res := tracesdk.WithResource(resource.NewSchemaless(
		semconv.ServiceNameKey.String(serviceName),
		semconv.ServiceInstanceIDKey.String(uuid),
		semconv.ProcessPIDKey.Int(os.Getpid()),
		attribute.String("environment", "development"),
		attribute.String("ip", ip.InternalIP()),
	))
	sampler := tracesdk.WithSampler(tracesdk.AlwaysSample())
	if c.Stdout != nil && c.Stdout.Enable && c.Stdout.Traces {
		var w io.Writer = os.Stdout
		if len(c.Stdout.TraceOutput) > 0 {
			if c.Stdout.TraceOutput == "no" {
				var n noOutput = 0
				w = &n
			} else {
				f, err := os.OpenFile(c.Stdout.TraceOutput, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModeAppend)
				if err != nil {
					return nil, err
				}
				w = f
			}
		}
		exp, err := stdouttrace.New(
			stdouttrace.WithWriter(w),
			stdouttrace.WithPrettyPrint(),
		)
		if err != nil {
			return nil, fmt.Errorf("creating stdout exporter: %w", err)
		}
		return tracesdk.NewTracerProvider(
			sampler,
			// Record information about this application in an Resource.
			res,
			// Always be sure to batch in production.
			tracesdk.WithSpanProcessor(tracesdk.NewBatchSpanProcessor(exp)),
		), nil
	} else if c.Otlp != nil && c.Otlp.Enable && c.Otlp.Traces {
		conn, err := grpc.DialContext(context.TODO(), c.Otlp.Endpoint,
			// Note the use of insecure transport here. TLS is recommended in production.
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock(),
		)
		if err != nil {
			return nil, err
		}
		exp, err := otlptracegrpc.New(context.TODO(), otlptracegrpc.WithGRPCConn(conn))
		if err != nil {
			return nil, err
		}
		if c.Otlp.Xray {
			idg := xray.NewIDGenerator()
			return tracesdk.NewTracerProvider(
				sampler,
				// Record information about this application in an Resource.
				res,
				// Always be sure to batch in production.
				tracesdk.WithSpanProcessor(tracesdk.NewBatchSpanProcessor(exp)),
				tracesdk.WithIDGenerator(idg),
			), nil
		} else {
			return tracesdk.NewTracerProvider(
				sampler,
				// Record information about this application in an Resource.
				res,
				// Always be sure to batch in production.
				tracesdk.WithSpanProcessor(tracesdk.NewBatchSpanProcessor(exp)),
			), nil
		}
	} else if c.Jaeger != nil && c.Jaeger.Enable && c.Jaeger.Traces {
		exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(c.Jaeger.Endpoint)))
		if err != nil {
			return nil, err
		}
		return tracesdk.NewTracerProvider(
			sampler,
			// Record information about this application in an Resource.
			res,
			// Always be sure to batch in production.
			tracesdk.WithSpanProcessor(tracesdk.NewBatchSpanProcessor(exp)),
		), nil
	} else {
		return nil, fmt.Errorf("not support trace")
	}
}

func newMeterProvider(c *Conf, serviceName string, uuid string) (sdk *metric.MeterProvider, err error) {
	var reader metric.Reader
	if c.Stdout != nil && c.Stdout.Enable && c.Stdout.Metrics {
		var w io.Writer = os.Stdout
		if len(c.Stdout.MetricOutput) > 0 {
			if c.Stdout.MetricOutput == "no" {
				var n noOutput = 0
				w = &n
			} else {
				f, err := os.OpenFile(c.Stdout.MetricOutput, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModeAppend)
				if err != nil {
					panic(err)
				}
				w = f
			}
		}
		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")
		exp, err := stdoutmetric.New(stdoutmetric.WithEncoder(enc))
		if err != nil {
			return nil, err
		}
		reader = metric.NewPeriodicReader(exp)
	} else if c.Otlp != nil && c.Otlp.Enable {
		ctx := context.Background()
		exp, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithEndpoint(c.Otlp.Endpoint), otlpmetricgrpc.WithInsecure())
		if err != nil {
			return nil, err
		}
		reader = metric.NewPeriodicReader(exp)
	} else if c.Prom != nil && c.Prom.Enable && c.Prom.Metrics {
		registry := prom.NewRegistry()
		reader, err = prometheus.New(prometheus.WithRegisterer(registry))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("not support metrics")
	}
	// Register the exporter with an SDK via a periodic reader.
	sdk = metric.NewMeterProvider(
		metric.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(serviceName),
		)),
		metric.WithReader(reader),
	)
	return sdk, nil
}
