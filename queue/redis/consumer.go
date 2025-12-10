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

func init() {
	queue.RegisterConsumer("redis", NewConsumer)
}

type Consumer struct {
	c       *queue.Conf
	handler queue.ConsumeHandler
	sub     *redis.Client
	wg      sync.WaitGroup
	ctx     context.Context
	cf      context.CancelFunc
}

func MustNewQueue(c *queue.Conf, handler queue.ConsumeHandler) queue.MessageQueue {
	q, err := NewConsumer(c, handler)
	if err != nil {
		log.Fatal(err)
	}
	return q
}

func NewConsumer(c *queue.Conf, handler queue.ConsumeHandler) (queue.MessageQueue, error) {
	ctx, cf := context.WithCancel(context.Background())
	opts := redis.Conf{
		Addrs: c.Brokers,
	}
	sub := redis.New(&opts)
	tracingSub := &Consumer{
		c:       c,
		handler: handler,
		sub:     sub,
		ctx:     ctx,
		cf:      cf,
	}
	tracingSub.wg.Add(1)
	go tracingSub.consumGroupTopic(ctx)
	log.Infof("start redis consumer, topic[%s]", tracingSub.c.Topic)
	return tracingSub, nil
}

func (c *Consumer) Name() string {
	return "redis"
}

func (c *Consumer) Subscribe(topic string, handler queue.ConsumeHandle, opts ...queue.SubscribeOption) (queue.Subscriber, error) {
	return nil, nil
}

func (c *Consumer) Close() error {
	c.cf()
	c.wg.Wait()
	c.sub.Close()
	log.Info("stop redis consumer. topic[%s]", c.c.Topic)
	return nil
}

func (c *Consumer) poll(ctx context.Context, timeoutMs time.Duration) (cctx context.Context, span trace.Span, err error) {
	serverTag, _ := os.Hostname()
	read, err := c.sub.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    c.c.Group,
		Consumer: serverTag,
		Streams:  []string{c.c.Topic, ">"},
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

func (c *Consumer) consumGroupTopic(ctx context.Context) {
	defer rescue.Recover(func() {
		c.wg.Done()
		log.Warnf("redis consumGroupTopic done")
	})
	serverTag, _ := os.Hostname()
	_, err := c.sub.XGroupCreateConsumer(c.ctx, c.c.Topic, c.c.Group, serverTag).Result()
	if err != nil && !strings.Contains(err.Error(), "BUSYGROUP Consumer Group name already exists") {
		log.Errorf("StreamGroupCreateConsumer[%v], err: %v", c.c.Topic, err)
		return
	}
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// ms
			cCtx, cf := context.WithTimeout(context.TODO(), 60*time.Second)
			cCtx, _, err := c.poll(cCtx, 100)
			if err != nil {
				log.Errorf("err: %v", err)
				cf()
				continue
			}
		}
	}
}
