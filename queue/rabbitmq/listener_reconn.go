package rabbitmq

import (
	"errors"
	"time"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/rescue"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (s *RabbitListener) reconnect() {
	defer rescue.Recover(func() {
		s.wg.Done()
	})
handleReconnect:
	var connCloseErr = make(chan *amqp.Error)
	var channelCloseErr = make(chan *amqp.Error)
	conn, err := amqp.Dial(s.c.Brokers[0])
	if err != nil {
		log.Error(err)
		return
	}
	conn.NotifyClose(connCloseErr)
	channel, err := conn.Channel()
	if err != nil {
		log.Error(err)
		return
	}
	channel.NotifyClose(channelCloseErr)
	for {
		select {
		case <-s.ctx.Done():
			channel.Close()
			conn.Close()
			log.Infof("[rabbitmq][listener] listener reconnect close")
			return
		case err := <-channelCloseErr:
			if err != nil && errors.Is(err, amqp.ErrClosed) {
				log.Errorf("[rabbitmq][listener] channel close notify: %v", err)
			} else if err != nil {
				log.Error(err)
			}
			time.Sleep(1 * time.Second)
			goto handleReconnect
		case err := <-connCloseErr:
			if err != nil && errors.Is(err, amqp.ErrClosed) {
				log.Errorf("[rabbitmq][listener] conn close notify: %v", err)
			} else if err != nil {
				log.Error(err)
			}
			time.Sleep(1 * time.Second)
			goto handleReconnect
		case t := <-s.topic:
			s.wg.Add(1)
			go s.run(s.ctx, channel, &ConsumerConf{
				Name:     t,
				Consumer: "xx",
			})
		}
	}
}
