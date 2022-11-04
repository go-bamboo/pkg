package rest

import (
	"context"
	"github.com/go-bamboo/pkg/middleware/logging"
	"github.com/go-bamboo/pkg/middleware/metadata"
	"github.com/go-bamboo/pkg/middleware/metrics/prometheus"
	"github.com/go-bamboo/pkg/tracing"

	"github.com/go-kratos/aegis/ratelimit/bbr"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go.opentelemetry.io/otel/propagation"
)

type Client = http.Client

func NewClient(c *Conf, opts ...Option) (*Client, error) {
	limiter := bbr.NewLimiter()
	middlewareChain := []middleware.Middleware{
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
	}
	c, err := http.NewClient(context.TODO(), http.WithEndpoint(""), http.WithMiddleware(middlewareChain...))
	if err != nil {
		return nil, err
	}
	return c, nil
}
