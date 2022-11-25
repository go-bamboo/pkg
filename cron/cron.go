package cron

import (
	"context"
	"go.uber.org/zap/zapcore"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/rescue"
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
		level zapcore.Level
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

func WithLogLevel(lvl zapcore.Level) Option {
	return func(alc *Cron) {
		alc.level = lvl
	}
}

// New 新cron
func New(options ...Option) (d *Cron) {
	s := &Cron{
		level: zapcore.DebugLevel,
	}
	for _, o := range options {
		o(s)
	}
	stdLogger := NewLogger(log.GetCore(), s.level)
	cr := cron.New(
		cron.WithSeconds(),
		cron.WithLogger(stdLogger),
	)
	for _, spec := range s.specs {
		cr.AddFunc(spec.spec, func() {
			defer rescue.Recover()
			if err := spec.handler(context.Background()); err != nil {
				log.ErrorStack(err)
			}
		})
	}
	s.cr = cr
	d = s
	return
}

func (c *Cron) Start() error {
	c.cr.Start()
	log.Infof("[cron] cron start")
	return nil
}

func (c *Cron) Stop() error {
	c.cr.Stop()
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
