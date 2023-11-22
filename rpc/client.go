package rpc

import (
	"context"
	"time"

	"github.com/go-bamboo/pkg/middleware/logging"
	"github.com/go-bamboo/pkg/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	// grpc options
	grpcInitialWindowSize     = 1 << 24
	grpcInitialConnWindowSize = 1 << 24
	grpcMaxSendMsgSize        = 1 << 24
	grpcMaxCallMsgSize        = 1 << 24
	grpcKeepAliveTime         = time.Second * 10
	grpcKeepAliveTimeout      = time.Second * 3
	grpcBackoffMaxDelay       = time.Second * 3
)

func NewClient(ctx context.Context, r registry.Discovery, appID string) (*ggrpc.ClientConn, error) {
	m := grpc.WithMiddleware(
		middleware.Chain(
			recovery.Recovery(),
			metadata.Client(metadata.WithPropagatedPrefix("x-md-global-", "X-Forwarded")),
			tracing.Client(),
			logging.Client(),
			validate.Validator(),
			//circuitbreaker.Client(),
		),
	)
	t := []grpc.ClientOption{
		m,
		grpc.WithEndpoint(appID),
		grpc.WithTimeout(time.Minute),
		grpc.WithDiscovery(r),
		grpc.WithOptions(
			ggrpc.WithInitialWindowSize(grpcInitialWindowSize),
			ggrpc.WithInitialConnWindowSize(grpcInitialConnWindowSize),
			ggrpc.WithDefaultCallOptions(ggrpc.MaxCallRecvMsgSize(grpcMaxCallMsgSize)),
			ggrpc.WithDefaultCallOptions(ggrpc.MaxCallSendMsgSize(grpcMaxSendMsgSize)),
			ggrpc.WithBackoffMaxDelay(grpcBackoffMaxDelay),
			ggrpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:                grpcKeepAliveTime,
				Timeout:             grpcKeepAliveTimeout,
				PermitWithoutStream: true,
			}),
		),
	}
	cc, err := grpc.DialInsecure(ctx, t...)
	if err != nil {
		return nil, err
	}
	return cc, nil
}
