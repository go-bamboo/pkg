package realip

import (
	"context"

	"github.com/go-bamboo/pkg/meta"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/tomasen/realip"
)

// Server is middleware server-side metadata.
func Server() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			_, err = meta.GetRealIP(ctx)
			if err != nil {
				if tr, ok := transport.FromServerContext(ctx); ok && tr.Kind() == transport.KindHTTP {
					req, ok := http.RequestFromServerContext(ctx)
					if ok {
						ip := realip.FromRequest(req)
						md, ok := metadata.FromServerContext(ctx)
						if ok {
							md.Set(meta.KeyRealIP, ip)
						}
					}
				}
			}
			return handler(ctx, req)
		}
	}
}
