package redis

import (
	"context"

	"github.com/go-bamboo/pkg/queue"
)

type Subscriber struct {
	topic  string
	cancel context.CancelFunc
	opts   queue.SubscribeOptions
}

func (s *Subscriber) Options() queue.SubscribeOptions {
	return s.opts
}

func (s *Subscriber) Topic() string {
	return s.topic
}

func (s *Subscriber) Unsubscribe(_ bool) error {
	if s.cancel != nil {
		s.cancel()
	}
	return nil
}
