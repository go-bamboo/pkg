package prometheus

import (
	metricsx "edu/pkg/stat/prom"
	prom "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
)

func Server() middleware.Middleware {
	return metrics.Server(
		metrics.WithSeconds(prom.NewHistogram(metricsx.MetricSeconds)),
		metrics.WithRequests(prom.NewCounter(metricsx.MetricRequests)),
	)
}
