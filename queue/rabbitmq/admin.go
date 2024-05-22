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
	conn, err := amqp.Dial(c.URL())
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

func (q *Admin) Bind(queueName string, routekey string, exchange, kind string, args amqp.Table) error {
	if err := q.channel.ExchangeDeclare(
		exchange,
		kind,
		true,  // 持久化
		false, // 自动删除
		false, // 非系统内部使用
		false, //
		args,
	); err != nil {
		return err
	}
	queue, err := q.channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		args,
	)
	if err != nil {
		return err
	}
	return q.channel.QueueBind(
		queue.Name,
		routekey,
		exchange,
		false,
		args,
	)
}
