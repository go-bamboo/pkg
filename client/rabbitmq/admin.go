package rabbitmq

import (
	"github.com/go-bamboo/pkg/log"
	"github.com/streadway/amqp"
)

type Admin struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func MustNewAdmin(c *RabbitConf) *Admin {
	conn, err := amqp.Dial(c.Address)
	if err != nil {
		log.Fatalf("failed to connect rabbitmq, error: %v", err)
	}
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open a channel, error: %v", err)
	}
	return &Admin{
		conn:    conn,
		channel: channel,
	}
}

func (q *Admin) DeclareExchange(conf *AdminExchangeConf, args amqp.Table) error {
	return q.channel.ExchangeDeclare(
		conf.Name,
		conf.Kind,
		conf.Durable,
		conf.AutoDelete,
		conf.Internal,
		conf.NoWait,
		args,
	)
}

func (q *Admin) DeclareQueue(conf *AdminQueueConf, args amqp.Table) error {
	_, err := q.channel.QueueDeclare(
		conf.Name,
		conf.Durable,
		conf.AutoDelete,
		conf.Exclusive,
		conf.NoWait,
		args,
	)
	return err
}

func (q *Admin) Bind(queueName string, routekey string, exchange string, notWait bool, args amqp.Table) error {
	return q.channel.QueueBind(
		queueName,
		routekey,
		exchange,
		notWait,
		args,
	)
}
