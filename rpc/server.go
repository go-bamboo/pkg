package rpc

import (
	"time"

	"github.com/go-bamboo/pkg/middleware/logging"
	"github.com/go-bamboo/pkg/middleware/metadata"
	"github.com/go-bamboo/pkg/middleware/metrics"
	"github.com/go-bamboo/pkg/middleware/tracing"
	"github.com/go-kratos/aegis/ratelimit/bbr"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type Server = grpc.Server

// NewServer new a gRPC server.
func NewServer(c *Conf) *Server {
	limiter := bbr.NewLimiter()
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			middleware.Chain(
				recovery.Recovery(),
				metadata.Server(),
				tracing.Server(),
				metrics.Server(),
				logging.Server(),
				ratelimit.Server(ratelimit.WithLimiter(limiter)),
				validate.Validator(),
			),
		),
		grpc.Options(
			ggrpc.KeepaliveParams(keepalive.ServerParameters{
				MaxConnectionIdle:     time.Second * 120,
				MaxConnectionAgeGrace: time.Second * 15,
				Time:                  time.Second * 30,
				Timeout:               time.Second * 10,
				MaxConnectionAge:      time.Hour * 4,
			}),
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
