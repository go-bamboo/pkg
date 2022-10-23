package main

import (
	"context"
	"github.com/go-bamboo/pkg/client/kafka"
	"github.com/go-bamboo/pkg/log"
	"os"
	"os/signal"
	"syscall"
)

type TestConsumerHandler struct {
}

func (c TestConsumerHandler) Consume(ctx context.Context, topic string, key, message []byte) error {
	log.Debug(message)
	return nil
}

func main() {
	cs, err := kafka.NewConsumer(&kafka.Conf{
		Brokers: []string{"192.168.1.13:9093", "192.168.1.160:9093", "192.168.1.203:9093"},
		Topic:   "test-can",
	}, &TestConsumerHandler{})
	if err != nil {
		log.Error(err)
	}
	cs.Start(context.TODO())
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	select {
	case s := <-c:
		log.Infof("get a signal [%s]\n", s.String())
	}
}
