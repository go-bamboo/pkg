package stdout

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/go-bamboo/pkg/net/ip"
	"github.com/go-bamboo/pkg/otel"
	otelx "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func init() {
	otel.Register("Stdout:Traces", NewTracerProvider)
	otel.Register("Stdout:Metrics", NewMeterProvider)
}

func NewTracerProvider(c *otel.Conf, serviceName string, uuid string) (err error) {
	res := tracesdk.WithResource(resource.NewSchemaless(
		semconv.ServiceNameKey.String(serviceName),
		semconv.ServiceInstanceIDKey.String(uuid),
		semconv.ProcessPIDKey.Int(os.Getpid()),
		attribute.String("environment", "development"),
		attribute.String("ip", ip.InternalIP()),
	))
	sampler := tracesdk.WithSampler(tracesdk.AlwaysSample())
	var w io.Writer = os.Stdout
	if len(c.Endpoint) > 0 {
		if c.Endpoint == "no" {
			var n noOutput = 0
			w = &n
		} else {
			f, err := os.OpenFile(c.Endpoint, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModeAppend)
			if err != nil {
				return err
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
		sampler,
		// Record information about this application in an Resource.
		res,
		// Always be sure to batch in production.
		tracesdk.WithSpanProcessor(tracesdk.NewBatchSpanProcessor(exp)),
	)
	otelx.SetTracerProvider(tp)
	return nil
}

func NewMeterProvider(c *otel.Conf, serviceName string, uuid string) (err error) {
	var reader metric.Reader
	var w io.Writer = os.Stdout
	if len(c.Endpoint) > 0 {
		if c.Endpoint == "no" {
			var n noOutput = 0
			w = &n
		} else {
			f, err := os.OpenFile(c.Endpoint, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModeAppend)
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
		return err
	}
	reader = metric.NewPeriodicReader(exp)

	sdk := metric.NewMeterProvider(
		metric.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(serviceName),
		)),
		metric.WithReader(reader),
	)
	otelx.SetMeterProvider(sdk)
	return nil
}
