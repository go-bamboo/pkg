package gormx

import (
	prom "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

type DBStats struct {
	MaxOpenConnections metrics.Gauge // Maximum number of open connections to the database.

	// Pool status
	OpenConnections metrics.Gauge // The number of established connections both in use and idle.
	InUse           metrics.Gauge // The number of connections currently in use.
	Idle            metrics.Gauge // The number of idle connections.

	// Counters
	WaitCount         metrics.Gauge // The total number of connections waited for.
	WaitDuration      metrics.Gauge // The total time blocked waiting for a new connection.
	MaxIdleClosed     metrics.Gauge // The total number of connections closed due to SetMaxIdleConns.
	MaxLifetimeClosed metrics.Gauge // The total number of connections closed due to SetConnMaxLifetime.
}

func newStats(labels map[string]string) *DBStats {
	stats := &DBStats{
		MaxOpenConnections: prom.NewGauge(prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name:        "gorm_dbstats_max_open_connections",
			Help:        "Maximum number of open connections to the database.",
			ConstLabels: labels,
		}, nil)),
		OpenConnections: prom.NewGauge(prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name:        "gorm_dbstats_open_connections",
			Help:        "The number of established connections both in use and idle.",
			ConstLabels: labels,
		}, nil)),
		InUse: prom.NewGauge(prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name:        "gorm_dbstats_in_use",
			Help:        "The number of connections currently in use.",
			ConstLabels: labels,
		}, nil)),
		Idle: prom.NewGauge(prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name:        "gorm_dbstats_idle",
			Help:        "The number of idle connections.",
			ConstLabels: labels,
		}, nil)),
		WaitCount: prom.NewGauge(prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name:        "gorm_dbstats_wait_count",
			Help:        "The total number of connections waited for.",
			ConstLabels: labels,
		}, nil)),
		WaitDuration: prom.NewGauge(prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name:        "gorm_dbstats_wait_duration",
			Help:        "The total time blocked waiting for a new connection.",
			ConstLabels: labels,
		}, nil)),
		MaxIdleClosed: prom.NewGauge(prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name:        "gorm_dbstats_max_idle_closed",
			Help:        "The total number of connections closed due to SetMaxIdleConns.",
			ConstLabels: labels,
		}, nil)),
		MaxLifetimeClosed: prom.NewGauge(prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name:        "gorm_dbstats_max_lifetime_closed",
			Help:        "The total number of connections closed due to SetConnMaxLifetime.",
			ConstLabels: labels,
		}, nil)),
	}

	//for _, collector := range stats.Collectors() {
	//	_ = prometheus.Register(collector)
	//}

	return stats
}
