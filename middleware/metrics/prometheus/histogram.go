package prometheus

import (
	"go.opentelemetry.io/otel/metric/global"

	"github.com/go-kratos/kratos/v2/metrics"
)

var _ metrics.Observer = (*histogram)(nil)

type histogram struct {
	lvs []string
}

// NewHistogram new a prometheus histogram and returns Histogram.
func NewHistogram() metrics.Observer {
	meterProvider := global.MeterProvider()
	meter := meterProvider.Meter("middleware.metrics")
	meter.AsyncInt64()
	return &histogram{}
}

func (h *histogram) With(lvs ...string) metrics.Observer {
	return &histogram{
		lvs: lvs,
	}
}

func (h *histogram) Observe(value float64) {
	//h.hv.WithLabelValues(h.lvs...).Observe(value)
}
