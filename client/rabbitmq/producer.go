package rabbitmq

import (
	"context"
	"github.com/emberfarkas/pkg/log"
	"github.com/emberfarkas/pkg/queue"
	"github.com/streadway/amqp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type (
	RabbitMqSender struct {
		c           ProducerConf
		ContentType string

		conn    *amqp.Connection
		channel *amqp.Channel

		tracer     trace.Tracer
		propagator propagation.TextMapPropagator
	}
)

func MustNewSender(c *ProducerConf) queue.Sender {
	sender := &RabbitMqSender{
		c:           *c,
		ContentType: c.ContentType,
	}
	if len(sender.ContentType) <= 0 {
		sender.ContentType = "text/plain"
	}
	conn, err := amqp.Dial(c.Rabbit.Address)
	if err != nil {
		log.Fatalf("failed to connect rabbitmq, error: %v", err)
	}
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open a channel, error: %v", err)
	}
	sender.conn = conn
	sender.channel = channel
	return sender
}

func (q *RabbitMqSender) Name() string {
	return ""
}

func (q *RabbitMqSender) Send(ctx context.Context, header map[string]interface{}, exchange string, routeKey string, msg []byte) error {
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
	log.Debugf("push exchange[%s], routeKey[%s] msg, header[%v]", exchange, routeKey, header)
	return nil
}

func (q *RabbitMqSender) Close() error {
	return q.conn.Close()
}
