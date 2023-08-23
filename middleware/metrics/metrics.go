package metrics

import (
	"context"
	"strconv"
	"time"

	"github.com/go-bamboo/pkg/middleware"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// Option is metrics option.
type Option func(*options)

// WithRequests with requests counter.
func WithRequests(c metric.Int64Counter) Option {
	return func(o *options) {
		o.requests = c
	}
}

// WithSeconds with seconds histogram.
func WithSeconds(c metric.Float64Histogram) Option {
	return func(o *options) {
		o.seconds = c
	}
}

type options struct {
	// counter: <client/server>_requests_code_total{kind, operation, code, reason}
	requests metric.Int64Counter
	// histogram: <client/server>_requests_seconds_bucket{kind, operation}
	seconds metric.Float64Histogram
}

// Server is middleware server-side metrics.
func Server(opts ...Option) middleware.Middleware {
	meter := otel.Meter("middleware.server")
	requests, _ := meter.Int64Counter("requests", metric.WithDescription("a requests counter"))
	seconds, _ := meter.Float64Histogram("seconds", metric.WithDescription("a seconds histogram"))
	op := options{
		requests: requests,
		seconds:  seconds,
	}
	for _, o := range opts {
		o(&op)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var (
				code      int
				reason    string
				kind      string
				operation string
			)
			startTime := time.Now()
			if info, ok := transport.FromServerContext(ctx); ok {
				kind = info.Kind().String()
				operation = info.Operation()
			}
			reply, err := handler(ctx, req)
			if se := errors.FromError(err); se != nil {
				code = int(se.Code)
				reason = se.Reason
			}
			if op.requests != nil {
				attrs := []attribute.KeyValue{
					attribute.Key("operation").String(operation),
					attribute.Key("code").String(strconv.Itoa(code)),
					attribute.Key("reason").String(reason),
				}
				op.requests.Add(ctx, 1, metric.WithAttributes(attrs...))
			}
			if op.seconds != nil {
				attrs := []attribute.KeyValue{
					attribute.Key("kind").String(kind),
					attribute.Key("operation").String(operation),
				}
				op.seconds.Record(ctx, time.Since(startTime).Seconds(), metric.WithAttributes(attrs...))
			}
			return reply, err
		}
	}
}

// Client is middleware client-side metrics.
func Client(opts ...Option) middleware.Middleware {
	meter := otel.Meter("middleware.client")
	requests, _ := meter.Int64Counter("requests", metric.WithDescription("a requests counter"))
	seconds, _ := meter.Float64Histogram("seconds", metric.WithDescription("a seconds histogram"))
	op := options{
		requests: requests,
		seconds:  seconds,
	}
	for _, o := range opts {
		o(&op)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var (
				code      int
				reason    string
				kind      string
				operation string
			)
			startTime := time.Now()
			if info, ok := transport.FromClientContext(ctx); ok {
				kind = info.Kind().String()
				operation = info.Operation()
			}
			reply, err := handler(ctx, req)
			if se := errors.FromError(err); se != nil {
				code = int(se.Code)
				reason = se.Reason
			}
			if op.requests != nil {
				attrs := []attribute.KeyValue{
					attribute.Key("operation").String(operation),
					attribute.Key("code").String(strconv.Itoa(code)),
					attribute.Key("reason").String(reason),
				}
				op.requests.Add(ctx, 1, metric.WithAttributes(attrs...))
			}
			if op.seconds != nil {
				attrs := []attribute.KeyValue{
					attribute.Key("kind").String(kind),
					attribute.Key("operation").String(operation),
				}
				op.seconds.Record(ctx, time.Since(startTime).Seconds(), metric.WithAttributes(attrs...))
			}
			return reply, err
		}
	}
}

func init() {
	middleware.Register("metrics-server", func(conf *middleware.Conf) (middleware.Middleware, error) {
		return Server(), nil
	})
	middleware.Register("metrics-client", func(conf *middleware.Conf) (middleware.Middleware, error) {
		return Server(), nil
	})
}
