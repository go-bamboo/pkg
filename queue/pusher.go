package queue

import "context"

// A Pusher interface wraps the method Push.
type Pusher interface {
	Name() string
	Push(ctx context.Context, key, value []byte) error
	Close() error
}

type Sender interface {
	Name() string
	Send(ctx context.Context, header map[string]interface{}, exchange string, routeKey string, msg []byte) error
	Close() error
}
