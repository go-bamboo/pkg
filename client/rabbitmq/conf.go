package rabbitmq

import (
	"fmt"
	"net/url"
)

func (c *RabbitConf) URL() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d%s", c.Username, url.QueryEscape(c.Password), c.Host, c.Port, c.VHost)
}
