package prom

import (
	"errors"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v3/cpu"
)

var (
	MetricSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "server",
		Subsystem: "requests",
		Name:      "duration_ms",
		Help:      "server requests duration(ms).",
		Buckets:   []float64{5, 10, 25, 50, 100, 250, 500, 1000},
	}, []string{"kind", "operation"})

	MetricRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "client",
		Subsystem: "requests",
		Name:      "code_total",
		Help:      "The total number of processed requests",
	}, []string{"kind", "operation", "code", "reason"})

	// 这个变量需要改代码
	MetricCPU = NewXCollector(XCollectorOpts{})
)

func init() {
	prometheus.MustRegister(MetricRequests, MetricSeconds)
}

type xCollector struct {
	collectFn       func(chan<- prometheus.Metric)
	pidFn           func() (int, error)
	reportErrors    bool
	cpuTotal        *prometheus.Desc
	openFDs, maxFDs *prometheus.Desc
	vsize, maxVsize *prometheus.Desc
	rss             *prometheus.Desc
	startTime       *prometheus.Desc
}

// ProcessCollectorOpts defines the behavior of a process metrics collector
// created with NewProcessCollector.
type XCollectorOpts struct {
	// PidFn returns the PID of the process the collector collects metrics
	// for. It is called upon each collection. By default, the PID of the
	// current process is used, as determined on construction time by
	// calling os.Getpid().
	PidFn func() (int, error)
	// If non-empty, each of the collected metrics is prefixed by the
	// provided string and an underscore ("_").
	Namespace string
	// If true, any error encountered during collection is reported as an
	// invalid metric (see NewInvalidMetric). Otherwise, errors are ignored
	// and the collected metrics will be incomplete. (Possibly, no metrics
	// will be collected at all.) While that's usually not desired, it is
	// appropriate for the common "mix-in" of process metrics, where process
	// metrics are nice to have, but failing to collect them should not
	// disrupt the collection of the remaining metrics.
	ReportErrors bool
}

// NewProcessCollector returns a collector which exports the current state of
// process metrics including CPU, memory and file descriptor usage as well as
// the process start time. The detailed behavior is defined by the provided
// ProcessCollectorOpts. The zero value of ProcessCollectorOpts creates a
// collector for the current process with an empty namespace string and no error
// reporting.
//
// The collector only works on operating systems with a Linux-style proc
// filesystem and on Microsoft Windows. On other operating systems, it will not
// collect any metrics.
func NewXCollector(opts XCollectorOpts) prometheus.Collector {
	ns := ""
	if len(opts.Namespace) > 0 {
		ns = opts.Namespace + "_"
	}

	c := &xCollector{
		reportErrors: opts.ReportErrors,
		cpuTotal: prometheus.NewDesc(
			ns+"process_cpu_seconds_total",
			"Total user and system CPU time spent in seconds.",
			nil, nil,
		),
		openFDs: prometheus.NewDesc(
			ns+"process_open_fds",
			"Number of open file descriptors.",
			nil, nil,
		),
		maxFDs: prometheus.NewDesc(
			ns+"process_max_fds",
			"Maximum number of open file descriptors.",
			nil, nil,
		),
		vsize: prometheus.NewDesc(
			ns+"process_virtual_memory_bytes",
			"Virtual memory size in bytes.",
			nil, nil,
		),
		maxVsize: prometheus.NewDesc(
			ns+"process_virtual_memory_max_bytes",
			"Maximum amount of virtual memory available in bytes.",
			nil, nil,
		),
		rss: prometheus.NewDesc(
			ns+"process_resident_memory_bytes",
			"Resident memory size in bytes.",
			nil, nil,
		),
		startTime: prometheus.NewDesc(
			ns+"process_start_time_seconds",
			"Start time of the process since unix epoch in seconds.",
			nil, nil,
		),
	}

	if opts.PidFn == nil {
		pid := os.Getpid()
		c.pidFn = func() (int, error) { return pid, nil }
	} else {
		c.pidFn = opts.PidFn
	}

	// Set up process metric collection if supported by the runtime.
	if canCollectProcess() {
		c.collectFn = c.processCollect
	} else {
		c.collectFn = func(ch chan<- prometheus.Metric) {
			c.reportError(ch, nil, errors.New("process metrics not supported on this platform"))
		}
	}

	return c
}

// Describe returns all descriptions of the collector.
func (c *xCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.cpuTotal
	ch <- c.openFDs
	ch <- c.maxFDs
	ch <- c.vsize
	ch <- c.maxVsize
	ch <- c.rss
	ch <- c.startTime
}

// Collect returns the current state of all metrics of the collector.
func (c *xCollector) Collect(ch chan<- prometheus.Metric) {
	c.collectFn(ch)
}

func (c *xCollector) reportError(ch chan<- prometheus.Metric, desc *prometheus.Desc, err error) {
	if !c.reportErrors {
		return
	}
	if desc == nil {
		desc = prometheus.NewInvalidDesc(err)
	}
	ch <- prometheus.NewInvalidMetric(desc, err)
}

func canCollectProcess() bool {
	return true
}

func (c *xCollector) processCollect(ch chan<- prometheus.Metric) {
	// pid, err := c.pidFn()
	// if err != nil {
	// 	c.reportError(ch, nil, err)
	// 	return
	// }

	// p, err := procfs.NewProc(pid)
	// if err != nil {
	// 	c.reportError(ch, nil, err)
	// 	return
	// }

	p, err := cpu.Percent(time.Second, true)
	if err != nil {
		c.reportError(ch, nil, err)
		return
	}
	for i := 0; i < len(p); i++ {
		ch <- prometheus.MustNewConstMetric(c.cpuTotal, prometheus.GaugeValue, p[i])
	}

	// if stat, err := p.Stat(); err == nil {
	// 	ch <- MustNewConstMetric(c.cpuTotal, CounterValue, stat.CPUTime())
	// 	ch <- MustNewConstMetric(c.vsize, GaugeValue, float64(stat.VirtualMemory()))
	// 	ch <- MustNewConstMetric(c.rss, GaugeValue, float64(stat.ResidentMemory()))
	// 	if startTime, err := stat.StartTime(); err == nil {
	// 		ch <- MustNewConstMetric(c.startTime, GaugeValue, startTime)
	// 	} else {
	// 		c.reportError(ch, c.startTime, err)
	// 	}
	// } else {
	// 	c.reportError(ch, nil, err)
	// }

	// if fds, err := p.FileDescriptorsLen(); err == nil {
	// 	ch <- MustNewConstMetric(c.openFDs, GaugeValue, float64(fds))
	// } else {
	// 	c.reportError(ch, c.openFDs, err)
	// }

	// if limits, err := p.Limits(); err == nil {
	// 	ch <- MustNewConstMetric(c.maxFDs, GaugeValue, float64(limits.OpenFiles))
	// 	ch <- MustNewConstMetric(c.maxVsize, GaugeValue, float64(limits.AddressSpace))
	// } else {
	// 	c.reportError(ch, nil, err)
	// }
}
