package rabbitmq

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/queue"
	"github.com/go-bamboo/pkg/rescue"
	"github.com/go-kratos/kratos/v2/errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type RabbitMqSender struct {
	c *queue.Conf
	//ContentType string

	isConnected     atomic.Bool
	isChannelOpen   atomic.Bool
	conn            *amqp.Connection
	channel         *amqp.Channel
	connCloseErr    chan *amqp.Error
	channelCloseErr chan *amqp.Error

	tracer     trace.Tracer
	propagator propagation.TextMapPropagator

	wg  sync.WaitGroup
	ctx context.Context
	cf  context.CancelFunc
}

func MustNewSender(c *queue.Conf) queue.Pusher {
	p, err := NewSender(c)
	if err != nil {
		log.Fatalf("new producer err: %v", err)
	}
	return p
}

func NewSender(c *queue.Conf) (queue.Pusher, error) {
	ctx, cf := context.WithCancel(context.TODO())
	sender := &RabbitMqSender{
		c: c,
		//ContentType:     c.ContentType,
		connCloseErr:    make(chan *amqp.Error),
		channelCloseErr: make(chan *amqp.Error),

		ctx: ctx,
		cf:  cf,
	}
	//if len(sender.ContentType) <= 0 {
	//	sender.ContentType = "text/plain"
	//}

	if err := sender.connect(); err != nil {
		return nil, err
	}
	if err := sender.open(); err != nil {
		return nil, err
	}
	sender.wg.Add(1)
	go sender.reconnect()
	return sender, nil
}

func (q *RabbitMqSender) Name() string {
	return "rabbitmq"
}

func (q *RabbitMqSender) Push(ctx context.Context, exchange string, routeKey []byte, msg []byte) error {
	if !q.isConnected.Load() {
		return ErrorDisconnect("%v", q.c.Brokers)
	}
	if !q.isChannelOpen.Load() {
		return ErrorChannelClosed("%v", q.c.Brokers)
	}
	header := map[string]interface{}{}
	body := Client(ctx, msg)
	err := q.channel.Publish(
		exchange,
		string(routeKey),
		false,
		false,
		amqp.Publishing{
			Headers:      header,
			ContentType:  "text/plain",
			Body:         body,
			DeliveryMode: amqp.Persistent,
		},
	)
	if err != nil {
		return wrapError(err)
	}
	log.Debugw("[rabbitmq][sender] push msg", "exchange", exchange, "routeKey", routeKey, "header", header)
	return nil
}

func (q *RabbitMqSender) PushWithPartition(ctx context.Context, exchange string, routeKey []byte, msg []byte, partition int32) error {
	if !q.isConnected.Load() {
		return ErrorDisconnect("%v", q.c.Brokers)
	}
	if !q.isChannelOpen.Load() {
		return ErrorChannelClosed("%v", q.c.Brokers)
	}
	header := map[string]interface{}{}
	body := Client(ctx, msg)
	err := q.channel.Publish(
		exchange,
		string(routeKey),
		false,
		false,
		amqp.Publishing{
			Headers:      header,
			ContentType:  "text/plain",
			Body:         body,
			DeliveryMode: amqp.Persistent,
		},
	)
	if err != nil {
		return wrapError(err)
	}
	log.Debugw("[rabbitmq][sender] push msg", "exchange", exchange, "routeKey", routeKey, "header", header)
	return nil
}

func (q *RabbitMqSender) Close() error {
	log.Infof("[rabbitmq][sender] stopping ...")
	q.cf()
	if q.isChannelOpen.Load() {
		if err := q.channel.Close(); err != nil {
			log.Error(err)
		}
	}
	if q.isConnected.Load() {
		if err := q.conn.Close(); err != nil {
			log.Error(err)
		}
	}
	q.wg.Wait()
	log.Infof("[rabbitmq][sender] stopping.")
	return nil
}

func (q *RabbitMqSender) connect() error {
	conn, err := amqp.Dial(q.c.Brokers[0])
	if err != nil {
		return err
	}
	conn.NotifyClose(q.connCloseErr)
	q.conn = conn
	q.isConnected.Store(true)
	log.Infof("[rabbitmq][sender] connected")
	return nil
}

func (q *RabbitMqSender) open() error {
	if !q.isConnected.Load() {
		return ErrorDisconnect("")
	}
	if q.isChannelOpen.Load() {
		return nil
	}
	channel, err := q.conn.Channel()
	if err != nil {
		return err
	}
	channel.NotifyClose(q.channelCloseErr)
	q.channel = channel
	q.isChannelOpen.Store(true)
	log.Infof("[rabbitmq][sender] channel open")
	return nil
}

func (q *RabbitMqSender) reconnect() {
	defer rescue.Recover(func() {
		q.wg.Done()
	})
	for {
		select {
		case <-q.ctx.Done():
			log.Infof("[rabbitmq][sender] sender reconnect close")
			return
		case err := <-q.channelCloseErr:
			if err != nil && errors.Is(err, amqp.ErrClosed) {
				log.Errorf("[rabbitmq][sender] channel close notify: %v", err)
				q.isChannelOpen.Store(false)
			} else if err != nil {
				log.Error(err)
			}
		case err := <-q.connCloseErr:
			if err != nil && errors.Is(err, amqp.ErrClosed) {
				log.Errorf("[rabbitmq][sender] conn close notify: %v", err)
				q.isConnected.Store(false)
				q.isChannelOpen.Store(false)
			}
		}
		if !q.isConnected.Load() {
			log.Infof("[rabbitmq][sender] Attempting to connect")
			if err := q.connect(); err != nil {
				log.Error(err)
			} else if err != nil {
				log.Error(err)
			}
		}
		if q.isConnected.Load() && !q.isChannelOpen.Load() {
			if err := q.open(); err != nil {
				log.Error(err)
			}
		}
		time.Sleep(time.Minute)
	}
}
