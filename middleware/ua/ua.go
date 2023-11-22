package ua

import (
	"context"

	"github.com/go-bamboo/pkg/meta"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

// Server is middleware server-side metadata.
func Server() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			md, ok := metadata.FromServerContext(ctx)
			if !ok {
				panic("metadata not found")
			}
			if ua := md.Get(meta.KeyUA); len(ua) <= 0 {
				if tr, ok := transport.FromServerContext(ctx); ok && tr.Kind() == transport.KindHTTP {
					header := tr.RequestHeader()
					md.Set(meta.KeyUA, header.Get("User-Agent"))
				}
			}
			return handler(ctx, req)
		}
	}
}
