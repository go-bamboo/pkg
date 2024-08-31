package elasticsearch

import (
	"net/http"

	otelext "github.com/go-bamboo/pkg/otel"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type EsTransportTracing struct {
	tracer     trace.Tracer
	kind       trace.SpanKind
	propagator propagation.TextMapPropagator
}

func NewEsTransportTracing() *EsTransportTracing {
	return &EsTransportTracing{
		tracer:     otel.Tracer("es"),
		kind:       trace.SpanKindClient,
		propagator: propagation.NewCompositeTextMapPropagator(otelext.Metadata{}, propagation.Baggage{}, otelext.TraceContext{}),
	}
}

func (h *EsTransportTracing) Perform(r *http.Request) (w *http.Response, err error) {
	//ctx := r.Context()
	//operation := "es:" + string(msg.Key)
	//ctx, span := p.tracer.Start(ctx, operation, nil)
	//err := p.pub.Produce(msg, deliveryChan)
	//p.tracer.End(ctx, span, nil, err)
	return
}
