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
	opts := redis.Conf{
		Addrs: c.Brokers,
	}
	pub := redis.New(&opts)
	tracingPub := &Producer{
		pub: pub,
	}
	return tracingPub, nil
}

func (p *Producer) Name() string {
	return "redis"
}

func (p *Producer) Push(ctx context.Context, topic string, key, value []byte) error {
	msg := redis.XAddArgs{
		Stream: topic,
		Approx: true,
		Values: value,
	}
	if err := p.pub.XAdd(ctx, &msg).Err(); err != nil {
		return err
	}
	return nil
}

func (p *Producer) PushWithPartition(ctx context.Context, topic string, key, value []byte, partition int32) error {
	msg := redis.XAddArgs{
		Stream: topic,
		Approx: true,
		Values: value,
	}
	if err := p.pub.XAdd(ctx, &msg).Err(); err != nil {
		return err
	}
	return nil
}

func (p *Producer) Close() error {
	return p.pub.Close()
}
