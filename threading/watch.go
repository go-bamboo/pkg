package threading

import (
	"context"
	"sync"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/rescue"
)

// A Watch is used to run given number of workers to process jobs.
type Watch struct {
	job     func(ctx context.Context, data interface{}) error
	workers int
	ch      chan interface{}

	wg  sync.WaitGroup
	ctx context.Context
	cf  context.CancelFunc
}

// NewWatch returns a Broadcast with given job and workers.
func NewWatch(job func(ctx context.Context, data interface{}) error, workers int) *Watch {
	ctx, cf := context.WithCancel(context.TODO())
	return &Watch{
		job:     job,
		workers: workers,
		ch:      make(chan interface{}, 1),

		ctx: ctx,
		cf:  cf,
	}
}

// Start starts a Broadcast.
func (b Watch) Start() error {
	for i := 0; i < b.workers; i++ {
		b.wg.Add(1)
		go func() {
			for {
				select {
				case <-b.ctx.Done():
					return
				case data := <-b.ch:
					b.run(data)
				}
			}
		}()
	}
	return nil
}

func (b Watch) Send(data interface{}) {
	b.ch <- data
}

func (b Watch) Stop() error {
	b.cf()
	b.wg.Wait()
	return nil
}

func (b Watch) run(data interface{}) {
	defer rescue.Recover(func() {
		b.wg.Done()
	})
	if err := b.job(b.ctx, data); err != nil {
		log.ErrorStack(err)
	}
}
