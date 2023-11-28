package rest

import (
	_ "net/http/pprof"

	"github.com/felixge/fgprof"
	"github.com/go-bamboo/pkg/api/status"
	"github.com/go-bamboo/pkg/middleware/logging"
	"github.com/go-bamboo/pkg/middleware/metrics"
	"github.com/go-bamboo/pkg/middleware/realip"
	"github.com/go-bamboo/pkg/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
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
		enc             http.EncodeResponseFunc
		ene             http.EncodeErrorFunc
	}
)

// WithMiddleware with sys client
func WithMiddleware(m ...middleware.Middleware) Option {
	return func(o *options) {
		o.middlewareChain = m
	}
}

func WithFilter(m ...http.FilterFunc) Option {
	return func(o *options) {
		o.filters = append(o.filters, m...)
	}
}

// ErrorEncoder with error encoder.
func ErrorEncoder(en http.EncodeErrorFunc) Option {
	return func(o *options) {
		o.ene = en
	}
}

// NewServer new a HTTP server.
func NewServer(c *Conf, opts ...Option) *Server {
	defaultOpts := &options{
		middlewareChain: []middleware.Middleware{
			recovery.Recovery(),
			metadata.Server(metadata.WithPropagatedPrefix("x-md-", "X-Forwarded")),
			realip.Server(), // 依赖metadata
			tracing.Server(),
			metrics.Server(),
			logging.Server(),
			validate.Validator(),
		},
		filters: make([]http.FilterFunc, 0),
		enc:     http.DefaultResponseEncoder,
		ene:     http.DefaultErrorEncoder,
	}
	for _, o := range opts {
		o(defaultOpts)
	}
	var serverOpts = []http.ServerOption{
		http.Filter(defaultOpts.filters...),
		http.Middleware(
			middleware.Chain(
				defaultOpts.middlewareChain...,
			),
		),
		http.ResponseEncoder(defaultOpts.enc),
		http.ErrorEncoder(defaultOpts.ene),
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
