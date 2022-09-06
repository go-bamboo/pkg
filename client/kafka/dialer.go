package kafka

import (
	"crypto/tls"
	"time"

	"github.com/segmentio/kafka-go"
)

type Dialer kafka.Dialer

func NewDialer(c *Conf) *Dialer {
	return &Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
		TLS:       &tls.Config{},
	}
}
