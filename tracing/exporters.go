package tracing

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/net/ip"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type Provider struct {
	tracer *tracesdk.TracerProvider
	meter  *metric.MeterProvider
}

func (o *Provider) Close() {
	if o.tracer != nil {
		o.tracer.Shutdown(context.TODO())
	}
	if o.meter != nil {
		o.meter.Shutdown(context.TODO())
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
	// Create the Jaeger exporter
	if c.Stdout.Enable {
		return stdoutProvider(c.Stdout, serviceName, uuid)
	}
	if c.Jaeger.Enable {
		return jaegerProvider(c.Jaeger, serviceName, uuid)
	}
	// 设置内部日志
	//otel.SetLogger()
	return nil, errors.New("not support")
}

func stdoutProvider(c *Stdout, serviceName string, uuid string) (*Provider, error) {
	var exporter Provider
	if c.Traces {
		exp, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
		if err != nil {
			return nil, fmt.Errorf("creating stdout exporter: %w", err)
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
		exporter.tracer = tp
	}
	if c.Metrics {
		// metric
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		exp, err := stdoutmetric.New(stdoutmetric.WithEncoder(enc))
		if err != nil {
			panic(err)
		}

		// Register the exporter with an SDK via a periodic reader.
		sdk := metric.NewMeterProvider(
			metric.WithResource(resource.NewSchemaless(
				semconv.ServiceNameKey.String("stdoutmetric-example"),
			)),
			metric.WithReader(metric.NewPeriodicReader(exp)),
		)
		global.SetMeterProvider(sdk)
		//meter.AsyncFloat64().
		exporter.meter = sdk
	}

	return &exporter, nil
}

func jaegerProvider(c *Jaeger, serviceName string, uuid string) (*Provider, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(c.Endpoint)))
	if err != nil {
		return nil, err
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
	return &Provider{tracer: tp}, nil
}
