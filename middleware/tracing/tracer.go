package tracing

import (
	"context"
	"fmt"

	otelext "github.com/go-bamboo/pkg/otel"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"
)

const Name = "jaeger"

// const Name = "kartos"

// Tracer is otel span tracer
type Tracer struct {
	tracer trace.Tracer
	kind   trace.SpanKind
	opt    *options
}

// NewTracer create tracer instance
func NewTracer(kind trace.SpanKind, opts ...Option) *Tracer {
	op := options{
		propagator: propagation.NewCompositeTextMapPropagator(otelext.Metadata{}, propagation.Baggage{}, otelext.TraceContext{}),
	}
	for _, o := range opts {
		o(&op)
	}
	if op.tracerProvider != nil {
		otel.SetTracerProvider(op.tracerProvider)
	}

	switch kind {
	case trace.SpanKindClient:
		return &Tracer{tracer: otel.Tracer(Name), kind: kind, opt: &op}
	case trace.SpanKindServer:
		return &Tracer{tracer: otel.Tracer(Name), kind: kind, opt: &op}
	case trace.SpanKindProducer:
		return &Tracer{tracer: otel.Tracer(Name), kind: kind, opt: &op}
	case trace.SpanKindConsumer:
		return &Tracer{tracer: otel.Tracer(Name), kind: kind, opt: &op}
	default:
		panic(fmt.Sprintf("unsupported span kind: %v", kind))
	}
}

// Start start otel span
func (t *Tracer) Start(ctx context.Context, operation string, carrier propagation.TextMapCarrier) (context.Context, trace.Span) {
	if t.kind == trace.SpanKindServer {
		ctx = t.opt.propagator.Extract(ctx, carrier)
	}
	if t.kind == trace.SpanKindConsumer {
		ctx = t.opt.propagator.Extract(ctx, carrier)
	}
	ctx, span := t.tracer.Start(ctx,
		operation,
		trace.WithSpanKind(t.kind),
	)
	if t.kind == trace.SpanKindClient {
		t.opt.propagator.Inject(ctx, carrier)
	}
	if t.kind == trace.SpanKindProducer {
		t.opt.propagator.Inject(ctx, carrier)
	}
	return ctx, span
}

// End finish otel span
func (t *Tracer) End(ctx context.Context, span trace.Span, m interface{}, err error) {
	if err != nil {
		span.RecordError(err)
		if e := errors.FromError(err); e != nil {
			span.SetAttributes(attribute.Key("rpc.status_code").Int64(int64(e.Code)))
		}
		span.SetStatus(codes.Error, err.Error())
	} else {
		span.SetStatus(codes.Ok, "OK")
	}

	if p, ok := m.(proto.Message); ok {
		if t.kind == trace.SpanKindServer {
			span.SetAttributes(attribute.Key("recv_msg.size").Int(proto.Size(p)))
		} else {
			span.SetAttributes(attribute.Key("send_msg.size").Int(proto.Size(p)))
		}
	}
	if p, ok := m.(*kafka.Message); ok {
		if t.kind == trace.SpanKindConsumer {
			span.SetAttributes(attribute.Key("sub_msg.size").Int(len(p.Value)))
		} else if t.kind == trace.SpanKindProducer {
			span.SetAttributes(attribute.Key("pub_msg.size").Int(len(p.Value)))
		}
	}
	span.End()
}
