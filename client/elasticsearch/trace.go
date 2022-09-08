package elasticsearch

import (
	"net/http"

	"github.com/emberfarkas/pkg/tracing"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type EsTransportTracing struct {
	tracer *tracing.Tracer
}

func NewEsTransportTracing() *EsTransportTracing {
	tracer := tracing.NewTracer(trace.SpanKindClient, tracing.WithPropagator(
		propagation.NewCompositeTextMapPropagator(tracing.Metadata{}, propagation.Baggage{}, tracing.TraceContext{}),
	))
	return &EsTransportTracing{
		tracer: tracer,
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
