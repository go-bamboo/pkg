package redis

import (
	"context"
	"github.com/go-bamboo/pkg/tracing"
	"github.com/go-redis/redis/extra/rediscmd/v8"
	"github.com/go-redis/redis/v8"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// RedisTracingHook redis的hook
type tracingHook struct {
	tracer     trace.Tracer
	propagator propagation.TextMapPropagator
}

var _ redis.Hook = (*tracingHook)(nil)

func NewRedisTracingHook() *tracingHook {
	return &tracingHook{
		tracer:     otel.Tracer("redis"),
		propagator: propagation.NewCompositeTextMapPropagator(tracing.Metadata{}, propagation.Baggage{}, tracing.TraceContext{}),
	}
}

func (h *tracingHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	operation := "redis:" + cmd.FullName()
	ctx, span := h.tracer.Start(ctx, operation, trace.WithSpanKind(trace.SpanKindClient))
	h.propagator.Inject(ctx, &RedisTextMapCarrier{})
	span.SetAttributes(
		attribute.String("db.system", "redis"),
		attribute.String("db.statement", rediscmd.CmdString(cmd)),
	)
	ctx = NewSpanContext(ctx, span)
	return ctx, nil
}

func (h *tracingHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	span, ok := SpanFromContext(ctx)
	if !ok {
		return ErrSpanLost("span is lost")
	}
	var err error
	if err = cmd.Err(); err != nil {
		recordError(ctx, span, err)
		cmd.SetErr(wrapRedisError(err))
	} else {
		span.SetAttributes(attribute.Key("redis.name").String(cmd.Name()))
	}
	return nil
}

func (h *tracingHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	var span trace.Span
	summary, cmdsString := rediscmd.CmdsString(cmds)
	operation := "redis:pipeline:" + summary
	ctx, span = h.tracer.Start(ctx, operation, trace.WithSpanKind(trace.SpanKindClient))
	span.SetAttributes(
		attribute.String("db.system", "redis"),
		attribute.Int("db.redis.num_cmd", len(cmds)),
		attribute.String("db.statement", cmdsString),
	)
	ctx = NewSpanContext(ctx, span)
	return ctx, nil
}

func (h *tracingHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	span, ok := SpanFromContext(ctx)
	if !ok {
		return ErrSpanLost("not found span")
	}
	var err error
	for _, cmd := range cmds {
		if err = cmd.Err(); err != nil {
			recordError(ctx, span, err)
			cmd.SetErr(wrapRedisError(err))
		} else {
			span.SetAttributes(attribute.Key("redis.name").String(cmd.Name()))
		}
	}
	return nil
}

func recordError(ctx context.Context, span trace.Span, err error) {
	if err != redis.Nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
}
