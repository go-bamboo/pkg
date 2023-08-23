package forwardcasbin

import (
	"context"
	"fmt"
	"net/url"

	syspb "github.com/go-bamboo/pkg/api/sys"
	"github.com/go-bamboo/pkg/meta"
	"github.com/go-bamboo/pkg/middleware"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport"
)

func ErrUnauthorizedPath(format string, a ...interface{}) error {
	return errors.Unauthorized("ErrUnauthorizedPath", fmt.Sprintf(format, a...))
}

func IsErrUnauthorizedPath(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "ErrUnauthorizedPath" && se.Code == 401
}

// Option is otel option.
type Option func(*options)

type options struct {
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
				dp := info.RequestHeader().Get(meta.KeyDP)
				uri, err1 := url.Parse(rawUri)
				if err1 != nil {
					err = err1
					return
				}
				x := req.(*syspb.CheckResourceRequest)
				x.Path = uri.Path
				x.Method = method
				reply, err = handler(ctx, x)
				info.ReplyHeader().Set(meta.KeyDP, dp)
				return
			}
			return handler(ctx, req)
		}
	}
}

func init() {
	middleware.Register("forwardcasbin", func(conf *middleware.Conf) (middleware.Middleware, error) {
		return Server(), nil
	})
}
