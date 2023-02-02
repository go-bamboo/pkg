package metadata

import (
	"context"
	"strings"

	"github.com/go-bamboo/pkg/meta"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

// Option is metadata option.
type Option func(*options)

type options struct {
	prefix []string
	md     metadata.Metadata
}

func (o *options) hasPrefix(key string) bool {
	k := strings.ToLower(key)
	for _, prefix := range o.prefix {
		if strings.HasPrefix(k, prefix) {
			return true
		}
	}
	return false
}

// WithConstants with constant metadata key value.
func WithConstants(md metadata.Metadata) Option {
	return func(o *options) {
		o.md = md
	}
}

// WithPropagatedPrefix with propagated key prefix.
func WithPropagatedPrefix(prefix ...string) Option {
	return func(o *options) {
		o.prefix = prefix
	}
}

// Server is middleware server-side metadata.
func Server(opts ...Option) middleware.Middleware {
	// X-Forwarded 兼容traffic
	options := &options{
		prefix: []string{"x-md-", "X-Forwarded"}, // x-md-global-, x-md-local
	}
	for _, o := range opts {
		o(options)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				md := options.md.Clone()
				header := tr.RequestHeader()
				for _, k := range header.Keys() {
					if options.hasPrefix(k) {
						md.Set(k, header.Get(k))
					}
					// 单独设置
					if k == "User-Agent" {
						md.Set(meta.KeyUA, header.Get(k))
					}
					if k == "X-Real-IP" {
						md.Set(meta.KeyRealIP, header.Get(k))
					}
					if k == "X-Forwarded-For" {
						ips := strings.Split(header.Get(k), ",")
						if len(ips) > 0 {
							md.Set(meta.KeyRealIP, ips[0])
						}
					}
				}
				ctx = metadata.NewServerContext(ctx, md)
			}
			return handler(ctx, req)
		}
	}
}

// Client is middleware client-side metadata.
func Client(opts ...Option) middleware.Middleware {
	options := &options{
		prefix: []string{"x-md-global-", "X-Forwarded"},
	}
	for _, o := range opts {
		o(options)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromClientContext(ctx); ok {
				header := tr.RequestHeader()
				// x-md-local-
				for k, v := range options.md {
					header.Set(k, v)
				}
				if md, ok := metadata.FromClientContext(ctx); ok {
					for k, v := range md {
						header.Set(k, v)
					}
				}
				// x-md-global-
				if md, ok := metadata.FromServerContext(ctx); ok {
					for k, v := range md {
						if options.hasPrefix(k) {
							header.Set(k, v)
						}
					}
				}
			}
			return handler(ctx, req)
		}
	}
}
