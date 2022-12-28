package threading

import (
	"context"
	"sync"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/rescue"
)

// A Watch is used to run given number of workers to process jobs.
type Watch[T any] struct {
	job     func(ctx context.Context, data T) error
	workers int
	ch      chan T

	wg sync.WaitGroup
}

// NewWatch returns a Watch with given job and workers.
func NewWatch[T any](job func(ctx context.Context, data T) error, workers int) *Watch[T] {
	return &Watch[T]{
		job:     job,
		workers: workers,
		ch:      make(chan T, 1),
	}
}

// Start starts a Broadcast.
func (b Watch[T]) Start() error {
	for i := 0; i < b.workers; i++ {
		b.wg.Add(1)
		go func() {
			for ch := range b.ch {
				b.run(ch)
			}
		}()
	}
	return nil
}

func (b Watch[T]) Send(data T) {
	b.ch <- data
}

func (b Watch[T]) Stop() error {
	close(b.ch)
	b.wg.Wait()
	return nil
}

func (b Watch[T]) run(data T) {
	defer rescue.Recover(func() {
		b.wg.Done()
	})
	if err := b.job(context.Background(), data); err != nil {
		log.ErrorStack(err)
	}
}
