package rest

import (
	"github.com/go-bamboo/pkg/middleware/realip"
	_ "net/http/pprof"

	"github.com/felixge/fgprof"
	"github.com/go-bamboo/pkg/api/status"
	"github.com/go-bamboo/pkg/middleware/logging"
	"github.com/go-bamboo/pkg/middleware/metadata"
	"github.com/go-bamboo/pkg/middleware/metrics"
	"github.com/go-bamboo/pkg/middleware/tracing"
	"github.com/go-kratos/aegis/ratelimit/bbr"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type (
	Server  = http.Server
	Option  func(*options)
	options struct {
		middlewareChain []middleware.Middleware
		filters         []http.FilterFunc
	}
)

// WithMiddleware with sys client
func WithMiddleware(m ...middleware.Middleware) Option {
	return func(o *options) {
		o.middlewareChain = append(o.middlewareChain, m...)
	}
}

func WithFilter(m ...http.FilterFunc) Option {
	return func(o *options) {
		o.filters = append(o.filters, m...)
	}
}

// NewServer new a HTTP server.
func NewServer(c *Conf, opts ...Option) *Server {
	limiter := bbr.NewLimiter()
	defaultOpts := &options{
		middlewareChain: []middleware.Middleware{
			recovery.Recovery(),
			metadata.Server(),
			realip.Server(),
			tracing.Server(),
			metrics.Server(),
			logging.Server(),
			ratelimit.Server(ratelimit.WithLimiter(limiter)),
		},
	}
	for _, o := range opts {
		o(defaultOpts)
	}
	defaultOpts.middlewareChain = append(defaultOpts.middlewareChain, validate.Validator())
	var serverOpts = []http.ServerOption{
		http.Filter(defaultOpts.filters...),
		http.Middleware(
			middleware.Chain(
				defaultOpts.middlewareChain...,
			),
		),
	}
	if c.Network != "" {
		serverOpts = append(serverOpts, http.Network(c.Network))
	}
	if c.Address != "" {
		serverOpts = append(serverOpts, http.Address(c.Address))
	}
	if c.Timeout != nil {
		serverOpts = append(serverOpts, http.Timeout(c.Timeout.AsDuration()))
	}
	httpSrv := http.NewServer(serverOpts...)
	status.RegisterStatusHTTPServer(httpSrv, status.NewStatusService())

	httpSrv.Handle("/debug/fgprof", fgprof.Handler())
	httpSrv.Handle("/metrics", promhttp.Handler())

	return httpSrv
}
