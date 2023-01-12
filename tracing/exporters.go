package tracing

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/net/ip"
	prom "github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
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
	var provider Provider
	// Create the Jaeger exporter
	if c.Stdout.Enable {
		stdoutProvider(&provider, c.Stdout, serviceName, uuid)
	}
	if c.Jaeger.Enable {
		jaegerProvider(&provider, c.Jaeger, serviceName, uuid)
	}
	if c.Otlp.Enable {
		otlpProvider(&provider, c.Otlp, serviceName, uuid)
	}
	if c.Prom.Enable {
		promProvider(&provider, c.Prom, serviceName, uuid)
	}
	// 设置内部日志
	//otel.SetLogger()
	return &provider, nil
}

type noOutput int

func (*noOutput) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func stdoutProvider(provider *Provider, c *Stdout, serviceName string, uuid string) error {
	if c.Traces {
		var w io.Writer = os.Stdout
		if len(c.TraceOutput) > 0 {
			if c.TraceOutput == "no" {
				var n noOutput = 0
				w = &n
			} else {
				f, err := os.OpenFile(c.TraceOutput, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModeAppend)
				if err != nil {
					panic(err)
				}
				w = f
			}
		}
		exp, err := stdouttrace.New(
			stdouttrace.WithWriter(w),
			stdouttrace.WithPrettyPrint(),
		)
		if err != nil {
			return fmt.Errorf("creating stdout exporter: %w", err)
		}
		tp := tracesdk.NewTracerProvider(
			tracesdk.WithSampler(tracesdk.AlwaysSample()),
			// Always be sure to batch in production.
			tracesdk.WithBatcher(exp),
			// Record information about this application in an Resource.
			tracesdk.WithResource(resource.NewSchemaless(
				semconv.ServiceNameKey.String(serviceName),
				semconv.ServiceInstanceIDKey.String(uuid),
				semconv.ProcessPIDKey.Int(os.Getpid()),
				attribute.String("environment", "development"),
				attribute.String("ip", ip.InternalIP()),
			)),
		)
		otel.SetTracerProvider(tp)
		provider.tracer = tp
	}
	if c.Metrics {
		// metric
		var w io.Writer = os.Stdout
		if len(c.MetricOutput) > 0 {
			if c.TraceOutput == "no" {
				var n noOutput = 0
				w = &n
			} else {
				f, err := os.OpenFile(c.MetricOutput, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModeAppend)
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
			panic(err)
		}
		// Register the exporter with an SDK via a periodic reader.
		sdk := metric.NewMeterProvider(
			metric.WithResource(resource.NewSchemaless(
				semconv.ServiceNameKey.String(serviceName),
			)),
			metric.WithReader(metric.NewPeriodicReader(exp)),
		)
		global.SetMeterProvider(sdk)
		//meter.AsyncFloat64().
		provider.meter = sdk
	}

	return nil
}

func jaegerProvider(provider *Provider, c *Jaeger, serviceName string, uuid string) error {
	if c.Traces && provider.tracer == nil {
		exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(c.Endpoint)))
		if err != nil {
			return err
		}
		tp := tracesdk.NewTracerProvider(
			tracesdk.WithSampler(tracesdk.AlwaysSample()),
			// Always be sure to batch in production.
			tracesdk.WithBatcher(exp),
			// Record information about this application in an Resource.
			tracesdk.WithResource(resource.NewSchemaless(
				semconv.ServiceNameKey.String(serviceName),
				semconv.ServiceInstanceIDKey.String(uuid),
				semconv.ProcessPIDKey.Int(os.Getpid()),
				attribute.String("environment", "development"),
				attribute.String("ip", ip.InternalIP()),
			)),
		)
		otel.SetTracerProvider(tp)
		provider.tracer = tp
	}

	return nil
}

func otlpProvider(provider *Provider, c *Otlp, serviceName string, uuid string) error {
	if c.Traces && provider.tracer == nil {
		client := otlptracegrpc.NewClient(otlptracegrpc.WithEndpoint(c.Endpoint), otlptracegrpc.WithInsecure())
		exporter, err := otlptrace.New(context.TODO(), client)
		if err != nil {
			return fmt.Errorf("creating OTLP trace exporter: %w", err)
		}
		tp := tracesdk.NewTracerProvider(
			tracesdk.WithSampler(tracesdk.AlwaysSample()),
			// Always be sure to batch in production.
			tracesdk.WithBatcher(exporter),
			// Record information about this application in an Resource.
			tracesdk.WithResource(resource.NewSchemaless(
				semconv.ServiceNameKey.String(serviceName),
				semconv.ServiceInstanceIDKey.String(uuid),
				semconv.ProcessPIDKey.Int(os.Getpid()),
				attribute.String("environment", "development"),
				attribute.String("ip", ip.InternalIP()),
			)),
		)
		otel.SetTracerProvider(tp)
		provider.tracer = tp
	}
	if c.Metrics && provider.meter == nil {
		ctx := context.Background()
		exp, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithEndpoint(c.Endpoint), otlpmetricgrpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		meterProvider := metric.NewMeterProvider(
			metric.WithResource(resource.NewSchemaless(
				semconv.ServiceNameKey.String(serviceName),
			)),
			metric.WithReader(metric.NewPeriodicReader(exp)),
		)
		global.SetMeterProvider(meterProvider)
		provider.meter = meterProvider
	}
	return nil
}

func promProvider(provider *Provider, c *Prometheus, serviceName string, uuid string) error {
	if c.Metrics && provider.meter == nil {
		registry := prom.NewRegistry()
		exporter, err := prometheus.New(prometheus.WithRegisterer(registry))
		if err != nil {
			panic(err)
		}
		meterProvider := metric.NewMeterProvider(
			metric.WithResource(resource.NewSchemaless(
				semconv.ServiceNameKey.String(serviceName),
			)),
			metric.WithReader(exporter),
		)
		global.SetMeterProvider(meterProvider)
		provider.meter = meterProvider
	}
	return nil
}
