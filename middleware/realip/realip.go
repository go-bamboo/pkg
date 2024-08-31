package realip

import (
	"context"

	"github.com/go-bamboo/pkg/meta"
	"github.com/go-bamboo/pkg/net/realip"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// Server is middleware server-side metadata.
func Server() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			md, ok := metadata.FromServerContext(ctx)
			if !ok {
				panic("metadata not found")
			}
			if ip := md.Get(meta.KeyRealIP); len(ip) <= 0 {
				if tr, ok := transport.FromServerContext(ctx); ok && tr.Kind() == transport.KindHTTP {
					if req, ok := http.RequestFromServerContext(ctx); ok {
						ip := realip.FromRequest(req)
						md.Set(meta.KeyRealIP, ip)
					}
				}
			}
			return handler(ctx, req)
		}
	}
}
