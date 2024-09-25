package queue

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport"
)

type (
	ConsumeHandle func(ctx context.Context, topic string, key, message []byte) error

	ConsumeHandler interface {
		Consume(ctx context.Context, topic string, key, message []byte) error
	}
)

type (
	// A MessageQueue interface represents a message queue.
	MessageQueue interface {
		transport.Server
		Name() string
	}
)
