package zipkin

import (
	"github.com/go-bamboo/pkg/net/ip"
	"github.com/go-bamboo/pkg/otel"
	otelx "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"os"
)

func init() {
	otel.Register("Zipkin:Traces", NewTracerProvider)
}

func NewTracerProvider(c *otel.Conf, serviceName string, uuid string, version string) (err error) {
	// exp
	exp, err := zipkin.New(c.Endpoint)
	if err != nil {
		return err
	}
	// res
	hostName, _ := os.Hostname()
	environment := c.Environment
	if len(environment) <= 0 {
		environment = "development"
	}
	res := tracesdk.WithResource(resource.NewSchemaless(
		semconv.ServiceNameKey.String(serviceName),
		semconv.ServiceInstanceIDKey.String(uuid),
		semconv.ServiceVersionKey.String(version),
		semconv.ProcessPIDKey.Int(os.Getpid()),
		semconv.DeploymentEnvironmentKey.String(environment),
		semconv.HostNameKey.String(hostName),
		attribute.String("ip", ip.InternalIP()),
	))
	sampler := tracesdk.WithSampler(tracesdk.AlwaysSample())
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
