package redis

import (
	"context"

	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type RedisTextMapCarrier struct {
}

var _ propagation.TextMapCarrier = &RedisTextMapCarrier{}

// Get returns the value associated with the passed key.
func (carrier *RedisTextMapCarrier) Get(key string) string {
	return ""
}

// Set stores the key-value pair.
func (carrier *RedisTextMapCarrier) Set(key string, value string) {
}

// Keys lists the keys stored in this carrier.
func (carrier *RedisTextMapCarrier) Keys() []string {
	return nil
}

func NewSpanContext(ctx context.Context, s trace.Span) context.Context {
	return context.WithValue(ctx, tracingHook{}, s)
}

// SpanFromContext returns the Transport value stored in ctx, if any.
func SpanFromContext(ctx context.Context) (s trace.Span, ok bool) {
	s, ok = ctx.Value(tracingHook{}).(trace.Span)
	return
}
