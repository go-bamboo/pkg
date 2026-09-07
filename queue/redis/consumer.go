package redis

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/queue"
	"github.com/go-bamboo/pkg/rescue"
	"github.com/go-bamboo/pkg/store/redis"
	stdredis "github.com/go-redis/redis/v8"
)

// Name is the name registered for redis
const Name = "redis"

var errBrokersEmpty = errors.New("redis queue: brokers is empty")

func init() {
	queue.RegisterConsumer(Name, NewConsumer)
	queue.RegisterPusher(Name, NewProducer)
}

type streamMessage struct {
	stream string
	id     string
	key    []byte
	value  []byte
}

type Consumer struct {
	c        *queue.Conf
	mu       sync.RWMutex
	handler  map[string]queue.ConsumeHandle
	sub      *redis.Client
	wg       sync.WaitGroup
	ctx      context.Context
	cf       context.CancelFunc
	group    string
	consumer string
}

func MustNewQueue(c *queue.Conf) queue.MessageQueue {
	q, err := NewConsumer(c)
	if err != nil {
		log.Fatal(err)
	}
	return q
}

func NewConsumer(c *queue.Conf) (queue.MessageQueue, error) {
	if c == nil || len(c.Brokers) == 0 {
		return nil, errBrokersEmpty
	}
	ctx, cf := context.WithCancel(context.Background())
	opts := redis.Conf{
		Addrs:        c.Brokers,
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
	}
	serverTag, _ := os.Hostname()
	if serverTag == "" {
		serverTag = Name
	}
	group := c.Group
	if group == "" {
		group = Name
	}
	q := &Consumer{
		c:        c,
		handler:  make(map[string]queue.ConsumeHandle),
		sub:      redis.New(&opts),
		ctx:      ctx,
		cf:       cf,
		group:    group,
		consumer: serverTag,
	}
	log.Infof("start redis consumer")
	return q, nil
}

func (c *Consumer) Name() string {
	return Name
}

func (c *Consumer) Subscribe(topic string, handler queue.ConsumeHandle, opts ...queue.SubscribeOption) (queue.Subscriber, error) {
	subOpts := queue.NewSubscribeOptions(opts...)
	c.mu.Lock()
	if _, ok := c.handler[topic]; ok {
		c.mu.Unlock()
		return nil, fmt.Errorf("redis queue: topic %s already subscribed", topic)
	}
	c.handler[topic] = handler
	c.mu.Unlock()

	ctx, cancel := context.WithCancel(c.ctx)
	c.wg.Add(1)
	go c.consumGroupTopic(ctx, topic)
	return &Subscriber{
		topic:  topic,
		cancel: cancel,
		opts:   subOpts,
	}, nil
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

func (c *Consumer) getHandler(topic string) queue.ConsumeHandle {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.handler[topic]
}

func (c *Consumer) poll(ctx context.Context, topic string, block time.Duration) ([]streamMessage, error) {
	read, err := c.sub.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    c.group,
		Consumer: c.consumer,
		Streams:  []string{topic, ">"},
		Count:    10,
		Block:    block,
	}).Result()
	if err != nil {
		if errors.Is(err, stdredis.Nil) || errors.Is(ctx.Err(), context.Canceled) || errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return nil, nil
		}
		return nil, err
	}
	out := make([]streamMessage, 0)
	for _, stream := range read {
		for _, message := range stream.Messages {
			key, value := payloadFromValues(message.Values)
			out = append(out, streamMessage{
				stream: stream.Stream,
				id:     message.ID,
				key:    key,
				value:  value,
			})
		}
	}
	return out, nil
}

func (c *Consumer) consumGroupTopic(ctx context.Context, topic string) {
	defer rescue.Recover(func() {
		c.wg.Done()
	})
	if c.c.Offset == "" {
		c.c.Offset = "0"
	}
	_, err := c.sub.XGroupCreateMkStream(ctx, topic, c.group, c.c.Offset).Result()
	if err != nil && !strings.Contains(err.Error(), "BUSYGROUP") {
		log.Errorf("StreamGroupCreateConsumer[%v], err: %v", topic, err)
		return
	}
	for {
		select {
		case <-ctx.Done():
			return
		default:
			pollCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
			msgs, err := c.poll(pollCtx, topic, 100*time.Millisecond)
			if err != nil {
				log.Errorf("err: %v", err)
				cancel()
				continue
			}
			handler := c.getHandler(topic)
			for _, msg := range msgs {
				if handler != nil {
					if herr := handler(pollCtx, topic, msg.key, msg.value); herr != nil {
						log.Errorf("redis consume handler: topic=%s id=%s err=%v", topic, msg.id, herr)
					}
				}
				if err := c.sub.XAck(ctx, msg.stream, c.group, msg.id).Err(); err != nil {
					log.Errorf("redis xack failed: stream=%s group=%s id=%s err=%v", msg.stream, c.group, msg.id, err)
				}
			}
			cancel()
		}
	}
}
