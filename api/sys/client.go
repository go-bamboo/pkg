package sys

import (
	"context"
	tracing2 "github.com/go-bamboo/pkg/tracing"
	"time"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/middleware/logging"
	"github.com/go-bamboo/pkg/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"go.opentelemetry.io/otel/propagation"
)

// AppID .
const appID = "discovery:///sys"

// NewSys new sso
func MustNew(ctx context.Context, timeout time.Duration, r registry.Discovery) (addressc SysClient) {
	m := grpc.WithMiddleware(
		middleware.Chain(
			recovery.Recovery(),
			metadata.Client(),
			tracing.Client(
				tracing.WithPropagator(
					propagation.NewCompositeTextMapPropagator(tracing2.Metadata{}, propagation.Baggage{}, tracing2.TraceContext{}),
				),
			),
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
