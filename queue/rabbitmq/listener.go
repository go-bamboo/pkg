package rabbitmq

import (
	"context"
	"sync"
	"time"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/queue"
	"github.com/go-bamboo/pkg/rescue"
	"github.com/go-kratos/kratos/v2/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

type (
	ConsumeHandle func(ctx context.Context, topic string, key, message []byte) error

	ConsumeHandler interface {
		Consume(ctx context.Context, topic string, key, message []byte) error
	}
	RabbitListener struct {
		c       *ListenerConf
		handler ConsumeHandler
		ctx     context.Context
		cf      context.CancelFunc
		wg      sync.WaitGroup
	}
)

func MustNewListener(c *ListenerConf, handler ConsumeHandler) queue.MessageQueue {
	listener, err := NewListener(c, handler)
	if err != nil {
		log.Fatal(err)
	}
	return listener
}

func NewListener(c *ListenerConf, handler ConsumeHandler) (consumer queue.MessageQueue, err error) {
	ctx, cf := context.WithCancel(context.Background())
	listener := &RabbitListener{
		c:       c,
		handler: handler,
		ctx:     ctx,
		cf:      cf,
	}
	return listener, nil
}

func (s *RabbitListener) Name() string {
	return "rabbitListener"
}

func (s *RabbitListener) Start(context.Context) error {
	for i := 0; i < len(s.c.Queues); i++ {
		q := s.c.Queues[i]
		s.wg.Add(1)
		go s.reconnect(q)
	}
	log.Infof("[rabbitmq][listener] start consume %d queue.", len(s.c.Queues))
	return nil
}

func (s *RabbitListener) Stop(context.Context) error {
	s.cf() // 停掉reconnect
	s.wg.Wait()
	log.Infof("[rabbitmq][listener] consumer stopping.")
	return nil
}

func (s *RabbitListener) run(ctx context.Context, channel *amqp.Channel, q *ConsumerConf) {
	defer rescue.Recover(func() {
		s.wg.Done()
	})
	for {
		select {
		case <-ctx.Done():
			log.Infof("[rabbitmq][listener] 结束消费队列[%s], exit.", q.Name)
			return
		default:
			// 获取消费通道
			log.Infof("[rabbitmq][listener] channel consume queue[%s]", q.Name)
			channel.Qos(1, 0, true) // 确保rabbitmq会一个一个发消息
			msgs, err := channel.Consume(
				q.Name,     // queue
				q.Consumer, // consumer
				false,      // auto-ack
				false,      // exclusive
				false,      // no-local
				true,       // no-wait
				nil,        // args
			)
			if nil != err {
				log.Errorf("[rabbitmq][listener] 获取队列[%s]的消费通道失败: %v, 结束消费队列", q.Name, err)
				return
			}
			log.Infof("[rabbitmq][listener] 开始处理[%s]消息", q.Name)
			for msg := range msgs {
				s.handleMsg(ctx, q.Name, &msg)
			}
		}
	}
}

func (s *RabbitListener) handleMsg(ctx context.Context, topic string, msg *amqp.Delivery) {
	// 当接收者消息处理失败的时候，
	// 比如网络问题导致的数据库连接失败，redis连接失败等等这种
	// 通过重试可以成功的操作，那么这个时候是需要重试的
	// 直到数据处理成功后再返回，然后才会回复rabbitmq ack
	if err := s.consume(ctx, topic, msg); err != nil {
		log.ErrorStack(err)
		if err = msg.Reject(false); err != nil {
			log.Error(err)
		}
		return
	}
	if err := msg.Ack(false); err != nil {
		log.Error(err)
	}
}

// consume 包裹一层，用来处理handler里面panic的error
func (s *RabbitListener) consume(ctx context.Context, topic string, msg *amqp.Delivery) error {
	c, cf, d, err := Server(msg.Body)
	if err != nil {
		return err
	}
	// 当遇到panic的error时候，现在处理是直接ack
	defer rescue.Recover(func() {
		cf()
	})
	return s.handler.Consume(c, topic, []byte(msg.RoutingKey), d)
}

func (s *RabbitListener) reconnect(c *ConsumerConf) {
	defer rescue.Recover(func() {
		s.wg.Done()
	})
handleReconnect:
	var connCloseErr chan *amqp.Error = make(chan *amqp.Error)
	var channelCloseErr chan *amqp.Error = make(chan *amqp.Error)
	conn, err := amqp.Dial(s.c.Rabbit.URL())
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
	s.wg.Add(1)
	go s.run(context.TODO(), channel, c)
	for {
		select {
		case <-s.ctx.Done():
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
		}
	}
}
