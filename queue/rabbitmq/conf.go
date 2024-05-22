package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
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
