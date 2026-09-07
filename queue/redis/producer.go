package redis

import (
	"context"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/queue"
	"github.com/go-bamboo/pkg/store/redis"
)

// Producer 生产者
type Producer struct {
	pub *redis.Client
}

func MustNewProducer(c *queue.Conf) queue.Pusher {
	pub, err := NewProducer(c)
	if err != nil {
		log.Fatal(err)
	}
	return pub
}

func NewProducer(c *queue.Conf) (queue.Pusher, error) {
	if c == nil || len(c.Brokers) == 0 {
		return nil, errBrokersEmpty
	}
	opts := redis.Conf{
		Addrs:        c.Brokers,
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
	}
	return &Producer{
		pub: redis.New(&opts),
	}, nil
}

func (p *Producer) Name() string {
	return Name
}

func (p *Producer) Push(ctx context.Context, topic string, key, value []byte) error {
	msg := redis.XAddArgs{
		Stream: topic,
		Values: map[string]interface{}{
			fieldKey:   string(key),
			fieldValue: string(value),
		},
	}
	return p.pub.XAdd(ctx, &msg).Err()
}

func (p *Producer) PushWithPartition(ctx context.Context, topic string, key, value []byte, _ int32) error {
	return p.Push(ctx, topic, key, value)
}

func (p *Producer) Close() error {
	return p.pub.Close()
}
