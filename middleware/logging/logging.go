package logging

import (
	"context"
	"fmt"
	"time"

	"github.com/go-bamboo/pkg/log"
	_ "github.com/go-bamboo/pkg/log/std"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

// Option is otel option.
type Option func(*options)

type options struct{}

// Server is an server logging middleware.
func Server(opts ...Option) middleware.Middleware {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var (
				code      int32
				reason    string
				kind      string
				operation string
			)
			startTime := time.Now()
			if info, ok := transport.FromServerContext(ctx); ok {
				kind = info.Kind().String()
				operation = info.Operation()
			}
			reply, err = handler(ctx, req)
			if se := errors.FromError(err); se != nil {
				code = se.Code
				reason = se.Reason
				log.Errorw(fmt.Sprintf("%+v", err),
					"kind", "server",
					"component", kind,
					"operation", operation,
					"args", extractArgs(req),
					"code", code,
					"reason", reason,
					"latency", time.Since(startTime).Seconds())
				return
			}
			log.Infow("",
				"kind", "server",
				"component", kind,
				"operation", operation,
				"args", extractArgs(req),
				"code", code,
				"reason", reason,
				"latency", time.Since(startTime).Seconds(),
			)
			return
		}
	}
}

// Client is an client logging middleware.
func Client(opts ...Option) middleware.Middleware {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var (
				code      int32
				reason    string
				kind      string
				operation string
			)
			startTime := time.Now()
			if info, ok := transport.FromClientContext(ctx); ok {
				kind = info.Kind().String()
				operation = info.Operation()
			}
			reply, err = handler(ctx, req)
			if se := errors.FromError(err); se != nil {
				code = se.Code
				reason = se.Reason
				log.Errorw(fmt.Sprintf("%+v", err),
					"kind", "client",
					"component", kind,
					"operation", operation,
					"args", extractArgs(req),
					"code", code,
					"reason", reason,
					"latency", time.Since(startTime).Seconds(),
				)
				return
			}
			log.Infow("",
				"kind", "client",
				"component", kind,
				"operation", operation,
				"args", extractArgs(req),
				"code", code,
				"reason", reason,
				"latency", time.Since(startTime).Seconds(),
			)
			return
		}
	}
}

// extractArgs returns the string of the req
func extractArgs(req interface{}) string {
	if stringer, ok := req.(fmt.Stringer); ok {
		return stringer.String()
	}
	return fmt.Sprintf("%+v", req)
}
