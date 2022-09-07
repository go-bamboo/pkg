package cron

import (
	"context"

	"github.com/emberfarkas/pkg/log"
	"github.com/emberfarkas/pkg/rescue"
	"github.com/robfig/cron/v3"
)

type (
	ConsumeHandle func(ctx context.Context) error

	Pair struct {
		spec    string
		handler ConsumeHandle
	}

	CronSrv struct {
		cr    *cron.Cron
		specs []Pair
	}

	Option func(alc *CronSrv)
)

func WithSpec(spec string, handler ConsumeHandle) Option {
	return func(alc *CronSrv) {
		alc.specs = append(alc.specs, Pair{spec: spec, handler: handler})
	}
}

// NewCron 新cron
func NewCron(options ...Option) (d *CronSrv) {
	stdLogger := log.NewZapLoggerEx(log.GetCore())
	cr := cron.New(
		cron.WithSeconds(),
		cron.WithLogger(stdLogger),
	)
	s := &CronSrv{
		cr: cr,
	}
	for _, o := range options {
		o(s)
	}
	for _, spec := range s.specs {
		cr.AddFunc(spec.spec, func() {
			defer rescue.Recover()
			if err := spec.handler(context.Background()); err != nil {
				log.ErrorStack(err)
			}
		})
	}
	d = s
	return
}

func (s *CronSrv) Start() error {
	s.cr.Start()
	log.Infof("[cron] cron start")
	return nil
}

func (s *CronSrv) Stop() error {
	s.cr.Stop()
	log.Infof("[cron] cron stop")
	return nil
}
