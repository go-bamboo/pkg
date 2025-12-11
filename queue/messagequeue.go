package queue

import (
	"context"
)

type (
	ConsumeHandle func(ctx context.Context, topic string, key, message []byte) error

	//ConsumeHandler interface {
	//	Consume(ctx context.Context, topic string, key, message []byte) error
	//}
)

// Subscriber .
type Subscriber interface {
	// Options .
	Options() SubscribeOptions

	// Topic .
	Topic() string

	// Unsubscribe .
	Unsubscribe(removeFromManager bool) error
}

// A MessageQueue interface represents a message queue.
type MessageQueue interface {
	Name() string
	Subscribe(topic string, handler ConsumeHandle, opts ...SubscribeOption) (Subscriber, error)
	Close() error
}
