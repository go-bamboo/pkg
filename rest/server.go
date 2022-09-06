package rest

import (
	_ "net/http/pprof"

	"edu/pkg/handler"
	"edu/pkg/middleware/logging"
	"edu/pkg/middleware/metadata"
	"edu/pkg/middleware/metrics/prometheus"
	"edu/pkg/tracing"

	"github.com/felixge/fgprof"
	"github.com/go-kratos/aegis/ratelimit/bbr"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/propagation"
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
			ratelimit.Server(ratelimit.WithLimiter(limiter)),
			metadata.Server(),
			tracing.Server(
				tracing.WithPropagator(
					propagation.NewCompositeTextMapPropagator(tracing.Metadata{}, propagation.Baggage{}, tracing.TraceContext{}),
				),
			),
			prometheus.Server(),
			logging.Server(),
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
	handler.RegisterStatusHTTPServer(httpSrv, handler.NewStatusService())

	httpSrv.Handle("/debug/fgprof", fgprof.Handler())
	httpSrv.Handle("/metrics", promhttp.Handler())

	return httpSrv
}
