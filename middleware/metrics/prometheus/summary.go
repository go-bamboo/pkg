package prometheus

import (
	"github.com/go-kratos/kratos/v2/metrics"
)

var _ metrics.Observer = (*summary)(nil)

type summary struct {
	lvs []string
}

// NewSummary new a prometheus summary and returns Histogram.
func NewSummary() metrics.Observer {
	return &summary{}
}

func (s *summary) With(lvs ...string) metrics.Observer {
	return &summary{
		lvs: lvs,
	}
}

func (s *summary) Observe(value float64) {
	//s.sv.WithLabelValues(s.lvs...).Observe(value)
}
