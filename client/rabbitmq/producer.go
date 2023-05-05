package rabbitmq

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/queue"
	"github.com/go-bamboo/pkg/rescue"
	"github.com/streadway/amqp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type (
	RabbitMqSender struct {
		c           *ProducerConf
		ContentType string

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
)

func MustNewSender(c *ProducerConf) queue.Sender {
	ctx, cf := context.WithCancel(context.TODO())
	sender := &RabbitMqSender{
		c:               c,
		ContentType:     c.ContentType,
		connCloseErr:    make(chan *amqp.Error),
		channelCloseErr: make(chan *amqp.Error),

		ctx: ctx,
		cf:  cf,
	}
	if len(sender.ContentType) <= 0 {
		sender.ContentType = "text/plain"
	}

	if err := sender.connect(); err != nil {
		log.Fatal(err)
	}
	if err := sender.open(); err != nil {
		log.Fatal(err)
	}
	sender.wg.Add(1)
	go sender.reconnect()
	return sender
}

func (q *RabbitMqSender) Name() string {
	return ""
}

func (q *RabbitMqSender) Send(ctx context.Context, header map[string]interface{}, exchange string, routeKey string, msg []byte) error {
	if !q.isConnected.Load() {
		return ErrorDisconnect("")
	}
	if !q.isChannelOpen.Load() {
		return ErrorChannelClosed("")
	}
	err := q.channel.Publish(
		exchange,
		routeKey,
		false,
		false,
		amqp.Publishing{
			Headers:      header,
			ContentType:  q.ContentType,
			Body:         msg,
			DeliveryMode: amqp.Persistent,
		},
	)
	if err != nil {
		return wrapError(err)
	}
	log.Debugw("[rabbitmq] push msg", "exchange", exchange, "routeKey", routeKey, "header", header)
	return nil
}

func (q *RabbitMqSender) Close() error {
	q.cf()
	q.wg.Wait()
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
	return nil
}

func (q *RabbitMqSender) connect() error {
	conn, err := amqp.Dial(q.c.Rabbit.Address)
	if err != nil {
		return err
	}
	conn.NotifyClose(q.connCloseErr)
	q.conn = conn
	q.isConnected.Store(true)
	log.Infof("[rabbitmq] connected")
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
	log.Infof("[rabbitmq] channel open")
	return nil
}

func (q *RabbitMqSender) reconnect() {
	defer rescue.Recover(func() {
		q.wg.Done()
	})
	for {
		if !q.isConnected.Load() {
			log.Infof("[rabbitmq] Attempting to connect")
			if err := q.connect(); err != nil {
				log.Error(err)
			}
		}
		if q.isConnected.Load() && !q.isChannelOpen.Load() {
			if err := q.open(); err != nil {
				log.Error(err)
			}
		}
		select {
		case <-q.ctx.Done():
			log.Infof("[rabbitmq] sender reconnect close")
			return
		case err := <-q.channelCloseErr:
			if err != nil {
				log.Errorf("[rabbitmq] channel close notify: %v", err)
				q.isChannelOpen.Store(false)
			}
		case err := <-q.connCloseErr:
			if err != nil {
				log.Errorf("[rabbitmq] conn close notify: %v", err)
				q.isConnected.Store(false)
			}
		}
		time.Sleep(time.Minute)
	}
}
