package prometheus

import (
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument/asyncint64"

	"github.com/go-kratos/kratos/v2/metrics"
)

var _ metrics.Counter = (*counter)(nil)

type counter struct {
	cv  asyncint64.Counter
	lvs []string
}

// NewCounter new a prometheus counter and returns Counter.
func NewCounter() metrics.Counter {
	meterProvider := global.MeterProvider()
	meter := meterProvider.Meter("gorm")
	cv, err := meter.AsyncInt64().Counter("xx")
	if err != nil {
		panic(err)
	}
	return &counter{
		cv: cv,
	}
}

func (c *counter) With(lvs ...string) metrics.Counter {
	return &counter{
		cv:  c.cv,
		lvs: lvs,
	}
}

func (c *counter) Inc() {
	//c.cv.WithLabelValues(c.lvs...).Inc()
}

func (c *counter) Add(delta float64) {
	//c.cv.WithLabelValues(c.lvs...).Add(delta)
}
