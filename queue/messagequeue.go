package queue

import (
	"context"
)

type (
	ConsumeHandle func(ctx context.Context, topic string, key, message []byte) error

	ConsumeHandler interface {
		Consume(ctx context.Context, topic string, key, message []byte) error
	}

	// A MessageQueue interface represents a message queue.
	MessageQueue interface {
		Name() string
		Start(context.Context) error
		Stop(context.Context) error
	}
)
