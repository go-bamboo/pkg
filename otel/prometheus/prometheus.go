package prometheus

import (
	"github.com/go-bamboo/pkg/otel"
	prom "github.com/prometheus/client_golang/prometheus"
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
	var reader metric.Reader
	registry := prom.NewRegistry()
	reader, err = prometheus.New(prometheus.WithRegisterer(registry))
	if err != nil {
		return err
	}
	// Register the exporter with an SDK via a periodic reader.
	sdk := metric.NewMeterProvider(
		metric.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(serviceName),
		)),
		metric.WithReader(reader),
	)
	otelx.SetMeterProvider(sdk)
	return nil
}
