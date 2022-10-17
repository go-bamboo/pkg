package threading

import (
	"context"
	"sync"
	"time"

	"github.com/emberfarkas/pkg/log"
	"github.com/emberfarkas/pkg/rescue"
)

// A PerpetualMotion is used to run given number of workers to process jobs.
type PerpetualMotion struct {
	job     func(ctx context.Context) (time.Duration, error)
	workers int

	wg  sync.WaitGroup
	ctx context.Context
	cf  context.CancelFunc

	Debug bool
}

// NewPerpetualMotion returns a NewPerpetualMotion with given job and workers.
func NewPerpetualMotion(job func(ctx context.Context) (time.Duration, error), workers int) PerpetualMotion {
	ctx, cf := context.WithCancel(context.TODO())
	return PerpetualMotion{
		job:     job,
		workers: workers,

		ctx: ctx,
		cf:  cf,
	}
}

// Start starts a WorkerGroup.
func (m PerpetualMotion) Start() error {
	for i := 0; i < m.workers; i++ {
		m.wg.Add(1)
		go func() {
			defer rescue.Recover(func() {
				log.Debug("[perpetual motion] done")
				m.wg.Done()
			})
			for {
				select {
				case <-m.ctx.Done():
					return
				default:
					m.run()
				}
			}
		}()
	}
	log.Infof("[perpetual motion] start")
	return nil
}

func (m PerpetualMotion) Stop() error {
	m.cf()
	m.wg.Wait()
	log.Infof("[perpetual motion] stop")
	return nil
}

func (m PerpetualMotion) run() {
	ctx, cf := context.WithTimeout(context.TODO(), time.Second*10)
	defer rescue.Recover(func() {
		cf()
	})
	delay, err := m.job(ctx)
	if err != nil {
		log.ErrorStack(err)
	}
	if delay > 0 {
		log.Debugf("sleep: %v", delay)
		time.Sleep(delay)
	}
}
