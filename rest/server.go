package rest

import (
	_ "net/http/pprof"

	"github.com/felixge/fgprof"
	"github.com/go-bamboo/pkg/api/status"
	"github.com/go-bamboo/pkg/middleware/logging"
	"github.com/go-bamboo/pkg/middleware/metadata"
	"github.com/go-bamboo/pkg/middleware/metrics"
	"github.com/go-bamboo/pkg/middleware/realip"
	"github.com/go-bamboo/pkg/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type (
	Server  = http.Server
	Option  func(*options)
	options struct {
		middlewareChain            []middleware.Middleware
		filters                    []http.FilterFunc
		enc                        http.EncodeResponseFunc
		ene                        http.EncodeErrorFunc
		loggingBalckListOperations []string
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

// ErrorEncoder with error encoder.
func ErrorEncoder(en http.EncodeErrorFunc) Option {
	return func(o *options) {
		o.ene = en
	}
}

func LoggingBalckListOperation(operation string) Option {
	return func(o *options) {
		o.loggingBalckListOperations = append(o.loggingBalckListOperations, operation)
	}
}

// NewServer new a HTTP server.
func NewServer(c *Conf, opts ...Option) *Server {
	defaultOpts := &options{
		middlewareChain: []middleware.Middleware{},
		filters:         make([]http.FilterFunc, 0),
		enc:             http.DefaultResponseEncoder,
		ene:             http.DefaultErrorEncoder,
	}
	for _, o := range opts {
		o(defaultOpts)
	}
	var loggingOpts []logging.Option
	for _, operation := range defaultOpts.loggingBalckListOperations {
		loggingOpts = append(loggingOpts, logging.WithBlackList(operation))
	}
	if len(defaultOpts.middlewareChain) <= 0 {
		defaultOpts.middlewareChain = append([]middleware.Middleware{
			recovery.Recovery(),
			realip.Server(),
			metadata.Server(),
			tracing.Server(),
			metrics.Server(),
			logging.Server(loggingOpts...),
			validate.Validator(),
		})
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
