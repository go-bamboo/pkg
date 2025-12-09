package rabbitmq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"net/url"
)

func (c *RabbitConf) URL() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d%s", c.Username, url.QueryEscape(c.Password), c.Host, c.Port, c.VHost)
}

var (
	ExchangeDirect  = amqp.ExchangeDirect
	ExchangeFanout  = amqp.ExchangeFanout
	ExchangeTopic   = amqp.ExchangeTopic
	ExchangeHeaders = amqp.ExchangeHeaders
)
