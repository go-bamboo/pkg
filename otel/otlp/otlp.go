package otlp

import (
	"context"
	"os"

	"github.com/go-bamboo/pkg/net/ip"
	"github.com/go-bamboo/pkg/otel"
	"go.opentelemetry.io/contrib/propagators/aws/xray"
	otelx "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	otel.Register("Otlp:Traces", NewTracerProvider)
	otel.Register("Otlp:Metrics", NewTracerProvider)
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
	conn, err := grpc.DialContext(context.TODO(), c.Endpoint,
		// Note the use of insecure transport here. TLS is recommended in production.
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return err
	}
	exp, err := otlptracegrpc.New(context.TODO(), otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return err
	}
	if c.Xray {
		idg := xray.NewIDGenerator()
		tp := tracesdk.NewTracerProvider(
			sampler,
			// Record information about this application in an Resource.
			res,
			// Always be sure to batch in production.
			tracesdk.WithSpanProcessor(tracesdk.NewBatchSpanProcessor(exp)),
			tracesdk.WithIDGenerator(idg),
		)
		otelx.SetTracerProvider(tp)
		return nil
	} else {
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
}

func NewMeterProvider(c *otel.Conf, serviceName string, uuid string) (err error) {
	var reader metric.Reader
	ctx := context.Background()
	exp, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithEndpoint(c.Endpoint), otlpmetricgrpc.WithInsecure())
	if err != nil {
		return err
	}
	reader = metric.NewPeriodicReader(exp)
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
