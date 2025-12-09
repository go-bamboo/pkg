package queue

import "context"

// A Pusher interface wraps the method Push.
type Pusher interface {
	Name() string
	Push(ctx context.Context, topic string, key, value []byte) error
	PushWithPartition(ctx context.Context, topic string, key, value []byte, partition int32) error
	Close() error
}
