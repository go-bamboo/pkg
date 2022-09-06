package forwardcasbin

import (
	"context"
	"fmt"
	"net/url"

	syspb "bls/api/sys"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

func ErrUnauthorizedPath(format string, a ...interface{}) error {
	return errors.Unauthorized("ErrUnauthorizedPath", fmt.Sprintf(format, a...))
}

func IsErrUnauthorizedPath(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "ErrUnauthorizedPath" && se.Code == 401
}

// Option is tracing option.
type Option func(*options)

type options struct {
	logger log.Logger
}

// WithLogger with sys client
func WithLogger(logger log.Logger) Option {
	return func(o *options) {
		o.logger = logger
	}
}

func Server(opts ...Option) middleware.Middleware {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if info, ok := transport.FromServerContext(ctx); ok {
				rawUri := info.RequestHeader().Get("X-Forwarded-Uri")
				method := info.RequestHeader().Get("X-Forwarded-Method")
				dp := info.RequestHeader().Get("x-md-global-dp")
				uri, err1 := url.Parse(rawUri)
				if err1 != nil {
					err = err1
					return
				}
				x := req.(*syspb.CheckResourceRequest)
				x.Path = uri.Path
				x.Method = method
				reply, err = handler(ctx, x)
				info.ReplyHeader().Set("x-md-global-dp", dp)
				return
			}
			return handler(ctx, req)
		}
	}
}
