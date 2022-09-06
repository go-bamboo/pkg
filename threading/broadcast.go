package threading

import (
	"bls/pkg/log"
	"bls/pkg/rescue"
	"context"
	"sync"
)

// A Broadcast is used to run given number of workers to process jobs.
type Broadcast struct {
	job     func(ctx context.Context, data interface{}) error
	workers int
	ch      chan interface{}

	wg  sync.WaitGroup
	ctx context.Context
	cf  context.CancelFunc
}

// NewBroadcast returns a Broadcast with given job and workers.
func NewBroadcast(job func(ctx context.Context, data interface{}) error, workers int) Broadcast {
	ctx, cf := context.WithCancel(context.TODO())
	return Broadcast{
		job:     job,
		workers: workers,
		ch:      make(chan interface{}),

		ctx: ctx,
		cf:  cf,
	}
}

// Start starts a Broadcast.
func (b Broadcast) Start() error {
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

func (b Broadcast) Send(data interface{}) {
	b.ch <- data
}

func (b Broadcast) Stop() error {
	b.cf()
	b.wg.Wait()
	return nil
}

func (b Broadcast) run(data interface{}) {
	defer rescue.Recover(func() {
		b.wg.Done()
	})
	if err := b.job(b.ctx, data); err != nil {
		log.ErrorStack(err)
	}
}
