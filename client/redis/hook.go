package redis

import (
	"context"

	"edu/pkg/tracing"

	"github.com/go-redis/redis/extra/rediscmd/v8"
	"github.com/go-redis/redis/v8"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

//RedisTracingHook redis的hook
type RedisTracingHook struct {
	tracer *tracing.Tracer
}

var _ redis.Hook = (*RedisTracingHook)(nil)

func NewRedisTracingHook() *RedisTracingHook {
	tracer := tracing.NewTracer(trace.SpanKindClient, tracing.WithPropagator(
		propagation.NewCompositeTextMapPropagator(tracing.Metadata{}, propagation.Baggage{}, tracing.TraceContext{}),
	))
	return &RedisTracingHook{
		tracer: tracer,
	}
}

func (h *RedisTracingHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	var span trace.Span
	operation := "redis:" + cmd.FullName()
	ctx, span = h.tracer.Start(ctx, operation, &RedisTextMapCarrier{})
	span.SetAttributes(
		attribute.String("db.system", "redis"),
		attribute.String("db.statement", rediscmd.CmdString(cmd)),
	)
	ctx = NewSpanContext(ctx, span)
	return ctx, nil
}

func (h *RedisTracingHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	span, ok := SpanFromContext(ctx)
	if !ok {
		return ErrSpanLost("span is lost")
	}
	var err error
	if err = cmd.Err(); err != nil {
		recordError(ctx, span, err)
		cmd.SetErr(wrapRedisError(err))
	}
	h.tracer.End(ctx, span, nil, err)
	return nil
}

func (h *RedisTracingHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	var span trace.Span
	summary, cmdsString := rediscmd.CmdsString(cmds)
	operation := "redis:pipeline:" + summary
	ctx, span = h.tracer.Start(ctx, operation, &RedisTextMapCarrier{})
	span.SetAttributes(
		attribute.String("db.system", "redis"),
		attribute.Int("db.redis.num_cmd", len(cmds)),
		attribute.String("db.statement", cmdsString),
	)
	ctx = NewSpanContext(ctx, span)
	return ctx, nil
}

func (h *RedisTracingHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	span, ok := SpanFromContext(ctx)
	if !ok {
		return ErrSpanLost("not found span")
	}
	var err error
	for _, cmd := range cmds {
		if err = cmd.Err(); err != nil {
			recordError(ctx, span, err)
			cmd.SetErr(wrapRedisError(err))
		}
	}
	h.tracer.End(ctx, span, nil, err)
	return nil
}

func recordError(ctx context.Context, span trace.Span, err error) {
	if err != redis.Nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
}
