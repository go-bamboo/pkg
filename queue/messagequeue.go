package queue

import (
	"github.com/go-kratos/kratos/v2/transport"
)

type (
	// A MessageQueue interface represents a message queue.
	MessageQueue interface {
		transport.Server
		Name() string
	}
)
