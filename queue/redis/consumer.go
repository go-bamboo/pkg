package redis

import (
	"context"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/queue"
	"github.com/go-bamboo/pkg/rescue"
	"github.com/go-bamboo/pkg/store/redis"
	"go.opentelemetry.io/otel/trace"
)

// Name is the name registered for redis
const Name = "redis"

func init() {
	queue.RegisterConsumer(Name, NewConsumer)
	queue.RegisterPusher(Name, NewProducer)
}

type Consumer struct {
	c       *queue.Conf
	handler map[string]queue.ConsumeHandle
	sub     *redis.Client
	wg      sync.WaitGroup
	ctx     context.Context
	cf      context.CancelFunc
}

func MustNewQueue(c *queue.Conf) queue.MessageQueue {
	q, err := NewConsumer(c)
	if err != nil {
		log.Fatal(err)
	}
	return q
}

func NewConsumer(c *queue.Conf) (queue.MessageQueue, error) {
	ctx, cf := context.WithCancel(context.Background())
	opts := redis.Conf{
		Addrs: c.Brokers,
	}
	sub := redis.New(&opts)
	tracingSub := &Consumer{
		c:       c,
		handler: make(map[string]queue.ConsumeHandle),
		sub:     sub,
		ctx:     ctx,
		cf:      cf,
	}

	log.Infof("start redis consumer")
	return tracingSub, nil
}

func (c *Consumer) Name() string {
	return "redis"
}

func (c *Consumer) Subscribe(topic string, handler queue.ConsumeHandle, opts ...queue.SubscribeOption) (queue.Subscriber, error) {
	c.handler[topic] = handler
	c.wg.Add(1)
	go c.consumGroupTopic(c.ctx, topic)
	return nil, nil
}

func (c *Consumer) Close() error {
	c.cf()
	c.wg.Wait()
	err := c.sub.Close()
	if err != nil {
		log.Errorf("stop redis consumer: %v", err)
		return err
	}
	log.Info("stop redis consumer")
	return nil
}

func (c *Consumer) poll(ctx context.Context, topic string, timeoutMs time.Duration) (cctx context.Context, span trace.Span, err error) {
	serverTag, _ := os.Hostname()
	read, err := c.sub.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    c.c.Group,
		Consumer: serverTag,
		Streams:  []string{topic, ">"},
		Block:    timeoutMs,
	}).Result()
	if err != nil {
		return ctx, nil, nil
	}
	for _, stream := range read {
		for _, message := range stream.Messages {
			if _, err := c.sub.XAck(ctx, stream.Stream, serverTag, message.ID).Result(); err != nil {
				log.Errorf("%v", err)
			}
		}
	}
	return
}

func (c *Consumer) consumGroupTopic(ctx context.Context, topic string) {
	defer rescue.Recover(func() {
		c.wg.Done()
		log.Warnf("redis consumGroupTopic done")
	})
	serverTag, _ := os.Hostname()
	_, err := c.sub.XGroupCreateConsumer(c.ctx, topic, c.c.Group, serverTag).Result()
	if err != nil && !strings.Contains(err.Error(), "BUSYGROUP Consumer Group name already exists") {
		log.Errorf("StreamGroupCreateConsumer[%v], err: %v", topic, err)
		return
	}
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// ms
			cCtx, cf := context.WithTimeout(context.TODO(), 60*time.Second)
			cCtx, _, err := c.poll(cCtx, topic, 100)
			if err != nil {
				log.Errorf("err: %v", err)
				cf()
				continue
			}
		}
	}
}
