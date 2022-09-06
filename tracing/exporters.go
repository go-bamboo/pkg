package tracing

import (
	"os"

	"edu/pkg/log"
	"edu/pkg/net/ip"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func MustTracerProvider(url string, serviceName string, uuid string) *tracesdk.TracerProvider {
	tp, err := TracerProvider(url, serviceName, uuid)
	if err != nil {
		log.Fatal(err)
	}
	return tp
}

// TracerProvider Get trace provider
func TracerProvider(url string, serviceName string, uuid string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
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
	return tp, nil
}
