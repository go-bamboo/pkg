package threading

import (
	"context"
	"sync"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/rescue"
)

// A Broadcast is used to run given number of workers to process jobs.
type Broadcast[T any] struct {
	job     func(ctx context.Context, data T) error
	workers int
	ch      chan T

	wg  sync.WaitGroup
	ctx context.Context
	cf  context.CancelFunc
}

// NewBroadcast returns a Broadcast with given job and workers.
func NewBroadcast[T any](job func(ctx context.Context, data T) error, workers int) *Broadcast[T] {
	ctx, cf := context.WithCancel(context.TODO())
	return &Broadcast[T]{
		job:     job,
		workers: workers,
		ch:      make(chan T),

		ctx: ctx,
		cf:  cf,
	}
}

// Start starts a Broadcast.
func (b Broadcast[T]) Start() error {
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

func (b Broadcast[T]) Send(data T) {
	b.ch <- data
}

func (b Broadcast[T]) Stop() error {
	b.cf()
	b.wg.Wait()
	return nil
}

func (b Broadcast[T]) run(data T) {
	defer rescue.Recover(func() {
		b.wg.Done()
	})
	if err := b.job(b.ctx, data); err != nil {
		log.ErrorStack(err)
	}
}
