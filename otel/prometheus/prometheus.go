package prometheus

import (
	"github.com/go-bamboo/pkg/otel"
	otelx "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func init() {
	otel.Register("Prometheus:Metrics", NewMeterProvider)
}

func NewMeterProvider(c *otel.Conf, serviceName string, uuid string) (err error) {
	exporter, err := prometheus.New()
	if err != nil {
		return err
	}
	// Register the exporter with an SDK via a periodic reader.
	provider := metric.NewMeterProvider(
		metric.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(serviceName),
			semconv.ServiceInstanceIDKey.String(uuid),
		)),
		metric.WithReader(exporter),
	)
	otelx.SetMeterProvider(provider)
	return nil
}
