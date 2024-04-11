package prometheus

import (
	"github.com/go-bamboo/pkg/net/ip"
	otelx "github.com/go-bamboo/pkg/otel"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"os"
)

func init() {
	otelx.Register("Prometheus:Metrics", NewMeterProvider)
}

func NewMeterProvider(c *otelx.Conf, serviceName string, uuid string, version string) (err error) {
	// exp
	exporter, err := prometheus.New()
	if err != nil {
		return err
	}
	// res
	hostName, _ := os.Hostname()
	environment := c.Environment
	if len(environment) <= 0 {
		environment = "development"
	}
	res := metric.WithResource(resource.NewSchemaless(
		semconv.ServiceNameKey.String(serviceName),
		semconv.ServiceInstanceIDKey.String(uuid),
		semconv.ServiceVersionKey.String(version),
		semconv.ProcessPIDKey.Int(os.Getpid()),
		semconv.DeploymentEnvironmentKey.String(environment),
		semconv.HostNameKey.String(hostName),
		attribute.String("ip", ip.InternalIP()),
	))
	// Register the exporter with an SDK via a periodic reader.
	provider := metric.NewMeterProvider(
		res,
		metric.WithReader(exporter),
	)
	otel.SetMeterProvider(provider)
	return nil
}
