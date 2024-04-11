package otlp

import (
	"context"
	"os"

	"github.com/go-bamboo/pkg/net/ip"
	otelx "github.com/go-bamboo/pkg/otel"
	"go.opentelemetry.io/contrib/propagators/aws/xray"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

func init() {
	otelx.Register("Otlp:Traces", NewTracerProvider)
	otelx.Register("Otlp:Metrics", NewMeterProvider)
}

func NewTracerProvider(c *otelx.Conf, serviceName string, uuid string, version string) (err error) {
	// exp
	grpcToken := c.GrpcToken
	if len(grpcToken) == 0 {
		grpcToken = os.Getenv("GRPC_TOKEN")
	}
	headers := map[string]string{"Authentication": grpcToken}
	exp, err := otlptracegrpc.New(context.TODO(),
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(c.Endpoint),
		otlptracegrpc.WithHeaders(headers),
		otlptracegrpc.WithDialOption(grpc.WithBlock()),
		otlptracegrpc.WithCompressor(gzip.Name))
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
	// sampler
	sampler := tracesdk.WithSampler(tracesdk.AlwaysSample())
	// xray
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
		otel.SetTracerProvider(tp)
		return nil
	} else {
		tp := tracesdk.NewTracerProvider(
			sampler,
			// Record information about this application in an Resource.
			res,
			// Always be sure to batch in production.
			tracesdk.WithSpanProcessor(tracesdk.NewBatchSpanProcessor(exp)),
		)
		otel.SetTracerProvider(tp)
		return nil
	}
}

func NewMeterProvider(c *otelx.Conf, serviceName string, uuid string, version string) (err error) {
	var reader metric.Reader
	ctx := context.Background()
	exp, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithEndpoint(c.Endpoint), otlpmetricgrpc.WithInsecure())
	if err != nil {
		return err
	}
	reader = metric.NewPeriodicReader(exp)
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
	sdk := metric.NewMeterProvider(
		res,
		metric.WithReader(reader),
	)
	otel.SetMeterProvider(sdk)
	return nil
}
