package jaeger

import (
	"fmt"
	"github.com/go-bamboo/pkg/net/ip"
	"github.com/go-bamboo/pkg/otel"
	otelx "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"os"
)

func init() {
	otel.Register(fmt.Sprintf("%s:%v", otel.ProviderType_Jaeger.String(), otel.Type_Traces.String()), NewTracerProvider)
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
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(c.Endpoint)))
	if err != nil {
		return err
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
