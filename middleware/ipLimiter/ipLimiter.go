package ipLimiter

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-bamboo/pkg/meta"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/patrickmn/go-cache"
	"golang.org/x/time/rate"
)

func ErrMaxLimit(format string, a ...interface{}) error {
	return errors.Unauthorized("ErrMaxLimit", fmt.Sprintf(format, a...))
}

// Option is otel option.
type Option func(*options)

type options struct {
	ips             sync.Map
	msec            int
	n               int
	enableBlackList bool
	blacklist       *cache.Cache
}

// WithSec with constant metadata key value.
func WithSec(sec int) Option {
	return func(o *options) {
		o.msec = sec
	}
}

func WithN(n int) Option {
	return func(o *options) {
		o.n = n
	}
}

func WithBlackList(enable bool) Option {
	return func(o *options) {
		o.enableBlackList = enable
	}
}

func Server(opts ...Option) middleware.Middleware {
	o := &options{
		msec:            200,
		n:               5,
		enableBlackList: false,
		blacklist:       cache.New(30*time.Second, 10*time.Minute),
	}
	for _, opt := range opts {
		opt(o)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			realIP, err := meta.GetRealIP(ctx)
			if err == nil {
				if o.enableBlackList {
					blocked, ok := o.blacklist.Get(realIP)
					if ok && blocked.(bool) {
						return nil, ErrMaxLimit("%v", realIP)
					}
				}
				var limiter *rate.Limiter
				dany, ok := o.ips.Load(realIP)
				if !ok {
					limiter = rate.NewLimiter(rate.Every(time.Duration(o.msec)*time.Millisecond), o.n)
					o.ips.Store(realIP, limiter)
				}
				limiter, ok = dany.(*rate.Limiter)
				if ok {
					if !limiter.Allow() {
						o.blacklist.SetDefault(realIP, true)
						return nil, ErrMaxLimit("%v", realIP)
					}
				}
			}
			return handler(ctx, req)
		}
	}
}
