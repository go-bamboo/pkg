package prometheus

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
)

func Server() middleware.Middleware {
	return metrics.Server(
		metrics.WithSeconds(NewHistogram()),
		metrics.WithRequests(NewCounter()),
	)
}

func Client() middleware.Middleware {
	return metrics.Client(
		metrics.WithSeconds(NewHistogram()),
		metrics.WithRequests(NewCounter()),
	)
}
