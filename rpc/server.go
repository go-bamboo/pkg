package rpc

import (
	"github.com/emberfarkas/pkg/middleware/logging"
	"github.com/emberfarkas/pkg/middleware/metadata"
	"github.com/emberfarkas/pkg/middleware/metrics/prometheus"
	"github.com/emberfarkas/pkg/tracing"
	"github.com/go-kratos/aegis/ratelimit/bbr"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"go.opentelemetry.io/otel/propagation"
)

type Server = grpc.Server

// NewServer new a gRPC server.
func NewServer(c *Conf) *Server {
	limiter := bbr.NewLimiter()
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			middleware.Chain(
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
				ratelimit.Server(),
				validate.Validator(),
			),
		),
	}
	if c.Network != "" {
		opts = append(opts, grpc.Network(c.Network))
	}
	if c.Address != "" {
		opts = append(opts, grpc.Address(c.Address))
	}
	if c.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)

	return srv
}
