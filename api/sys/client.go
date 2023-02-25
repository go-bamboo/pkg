package sys

import (
	"context"
	"time"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/middleware/logging"
	"github.com/go-bamboo/pkg/middleware/metrics"
	"github.com/go-bamboo/pkg/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// AppID .
const appID = "discovery:///sys"

// MustNew new sys
func MustNew(ctx context.Context, timeout time.Duration, r registry.Discovery) (addressc SysClient) {
	m := grpc.WithMiddleware(
		middleware.Chain(
			recovery.Recovery(),
			metadata.Client(),
			tracing.Client(),
			metrics.Client(),
			logging.Client(),
			validate.Validator(),
			//circuitbreaker.Client(),
		),
	)
	t := []grpc.ClientOption{
		m,
		grpc.WithEndpoint(appID),
		grpc.WithTimeout(timeout),
		grpc.WithDiscovery(r),
	}
	cc, err := grpc.DialInsecure(ctx, t...)
	if err != nil {
		log.Fatal(err)
	}
	return NewSysClient(cc)
}
