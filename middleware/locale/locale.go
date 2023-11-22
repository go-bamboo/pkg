package locale

import (
	"context"
	"github.com/go-bamboo/pkg/meta"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"strings"
)

// Server is middleware server-side metadata.
func Server() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			md, ok := metadata.FromServerContext(ctx)
			if !ok {
				panic("metadata not found")
			}
			if lo := md.Get(meta.KeyLocale); len(lo) <= 0 {
				if tr, ok := transport.FromServerContext(ctx); ok && tr.Kind() == transport.KindHTTP {
					header := tr.RequestHeader()
					var locale string
					locale1 := header.Get("Locale")
					locale2 := header.Get("locale")
					if len(locale1) > 0 {
						locale = locale1
					}
					if len(locale2) > 0 {
						locale = locale2
					}
					locale = strings.ToUpper(locale)
					if strings.Contains(locale, "EN") {
						locale = "EN"
					}
					if strings.Contains(locale, "CN") {
						locale = "CN"
					}
					if locale != "EN" && locale != "CN" {
						locale = "EN"
					}
					md.Set(meta.KeyLocale, locale)
				}
			}
			return handler(ctx, req)
		}
	}
}
