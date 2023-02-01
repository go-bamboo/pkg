package tracing

import (
	"context"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/metadata"
	"go.opentelemetry.io/otel/propagation"
)

const ServiceHeader = "x-md-service-name"

// Metadata is tracing metadata propagator
type Metadata struct{}

var _ propagation.TextMapPropagator = Metadata{}

// Inject sets metadata key-values from ctx into the carrier.
func (b Metadata) Inject(ctx context.Context, carrier propagation.TextMapCarrier) {
	app, ok := kratos.FromContext(ctx)
	if ok {
		carrier.Set(ServiceHeader, app.Name())
	}
}

// Extract returns a copy of parent with the metadata from the carrier added.
func (b Metadata) Extract(parent context.Context, carrier propagation.TextMapCarrier) context.Context {
	name := carrier.Get(ServiceHeader)
	if name != "" {
		if md, ok := metadata.FromServerContext(parent); ok {
			md.Set(ServiceHeader, name)
		} else {
			md := metadata.New()
			md.Set(ServiceHeader, name)
			parent = metadata.NewServerContext(parent, md)
		}
	}

	return parent
}

// Fields returns the keys who's values are set with Inject.
func (b Metadata) Fields() []string {
	return []string{ServiceHeader}
}
