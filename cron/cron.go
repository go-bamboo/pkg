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

	Cron struct {
		cr    *cron.Cron
		specs []Pair
	}

	Option func(alc *Cron)

	EntryID = cron.EntryID

	Job = cron.Job
)

func WithSpec(spec string, handler ConsumeHandle) Option {
	return func(alc *Cron) {
		alc.specs = append(alc.specs, Pair{spec: spec, handler: handler})
	}
}

// New 新cron
func New(options ...Option) (d *Cron) {
	stdLogger := log.NewZapLoggerEx(log.GetCore())
	cr := cron.New(
		cron.WithSeconds(),
		cron.WithLogger(stdLogger),
	)
	s := &Cron{
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

func (s *Cron) Start() error {
	s.cr.Start()
	log.Infof("[cron] cron start")
	return nil
}

func (s *Cron) Stop() error {
	s.cr.Stop()
	log.Infof("[cron] cron stop")
	return nil
}

func (c *Cron) AddFunc(spec string, cmd func()) (EntryID, error) {
	return c.cr.AddFunc(spec, cmd)
}

func (c *Cron) AddJob(spec string, cmd Job) (EntryID, error) {
	return c.cr.AddJob(spec, cmd)
}

func (c *Cron) Remove(id EntryID) {
	c.cr.Remove(id)
}
