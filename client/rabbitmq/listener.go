package rabbitmq

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/queue"
	"github.com/go-bamboo/pkg/rescue"
	"github.com/streadway/amqp"
)

type (
	ConsumeHandle func(ctx context.Context, topic string, key, message []byte) error

	ConsumeHandler interface {
		Consume(ctx context.Context, topic string, key, message []byte) error
	}

	RabbitListener struct {
		c               *ListenerConf
		isConnected     atomic.Bool
		isChannelOpen   atomic.Bool
		conn            *amqp.Connection
		channel         *amqp.Channel
		connCloseErr    chan *amqp.Error
		channelCloseErr chan *amqp.Error
		handler         ConsumeHandler

		ctx context.Context
		cf  context.CancelFunc
		wg  sync.WaitGroup
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
		c:               c,
		connCloseErr:    make(chan *amqp.Error),
		channelCloseErr: make(chan *amqp.Error),
		handler:         handler,

		ctx: ctx,
		cf:  cf,
	}
	if err := listener.connect(); err != nil {
		return nil, err
	}
	if err := listener.open(); err != nil {
		return nil, err
	}
	return
}

func (s *RabbitListener) Name() string {
	return "rabbitListener"
}

func (s *RabbitListener) Start(context.Context) error {
	s.wg.Add(1)
	go s.reconnect()
	for i := 0; i < len(s.c.Queues); i++ {
		q := s.c.Queues[i]
		s.wg.Add(1)
		go s.run(s.ctx, q)
	}
	log.Infof("[rabbitmq] start consume %d queue.", len(s.c.Queues))
	return nil
}

func (s *RabbitListener) Stop(context.Context) error {
	if s.isChannelOpen.Load() {
		if err := s.channel.Close(); err != nil {
			log.Error(err)
		}
	}
	if s.isConnected.Load() {
		if err := s.conn.Close(); err != nil {
			log.Error(err)
		}
	}
	s.cf()
	s.wg.Wait()
	if s.isChannelOpen.Load() {
		if err := s.channel.Close(); err != nil {
			log.Error(err)
		}
	}
	if s.isConnected.Load() {
		if err := s.conn.Close(); err != nil {
			log.Error(err)
		}
	}
	log.Infof("[rabbitmq] stop consumer.")
	return nil
}

func (s *RabbitListener) run(ctx context.Context, q *ConsumerConf) {
	defer rescue.Recover(func() {
		s.wg.Done()
	})
	for {
		select {
		case <-ctx.Done():
			log.Warnf("[rabbitmq] 结束消费队列[%s], exit.", q.Name)
			return
		default:
			if !s.isConnected.Load() {
				log.Warnf("rabbitmq disconnect")
				time.Sleep(30 * time.Second)
				continue
			}
			if !s.isChannelOpen.Load() {
				log.Warnf("rabbitmq channel not open")
				time.Sleep(30 * time.Second)
				continue
			}
			// 获取消费通道
			s.channel.Qos(1, 0, true) // 确保rabbitmq会一个一个发消息
			log.Infof("[rabbitmq] channel consume queue[%s]", q.Name)
			msgs, err := s.channel.Consume(
				q.Name,     // queue
				q.Consumer, // consumer
				false,      // auto-ack
				false,      // exclusive
				false,      // no-local
				false,      // no-wait
				nil,        // args
			)
			if nil != err {
				log.Errorf("[rabbitmq] 获取队列[%s]的消费通道失败: %v", q.Name, err)
				return
			}
			log.Infof("[rabbitmq] 开始处理[%s]消息", q.Name)
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
	c, cf := context.WithTimeout(context.TODO(), time.Second*10)
	// 当遇到panic的error时候，现在处理是直接ack
	defer rescue.Recover(func() {
		cf()
	})
	return s.handler.Consume(c, topic, []byte(msg.RoutingKey), msg.Body)
}

func (s *RabbitListener) connect() error {
	conn, err := amqp.Dial(s.c.Rabbit.Address)
	if err != nil {
		return err
	}
	conn.NotifyClose(s.connCloseErr)
	s.conn = conn
	s.isConnected.Store(true)
	log.Infof("rabbitmq connected")
	return nil
}

func (s *RabbitListener) open() error {
	if !s.isConnected.Load() {
		return ErrorDisconnect("")
	}
	if s.isChannelOpen.Load() {
		return nil
	}
	channel, err := s.conn.Channel()
	if err != nil {
		return err
	}
	channel.NotifyClose(s.channelCloseErr)
	s.channel = channel
	s.isChannelOpen.Store(true)
	log.Infof("rabbitmq channel open")
	return nil
}

func (s *RabbitListener) reconnect() {
	defer rescue.Recover(func() {
		s.wg.Done()
	})
	for {
		if !s.isConnected.Load() {
			log.Infof("Attempting to connect")
			if err := s.connect(); err != nil {
				log.Error(err)
			}
		}
		if s.isConnected.Load() && !s.isChannelOpen.Load() {
			if err := s.open(); err != nil {
				log.Error(err)
			}
		}
		select {
		case <-s.ctx.Done():
			log.Infof("rabbitmq sender close")
			return
		case err := <-s.channelCloseErr:
			if err != nil {
				log.Errorf("channel close notify: %v", err)
				s.isChannelOpen.Store(false)
			}
		case err := <-s.connCloseErr:
			if err != nil {
				log.Errorf("conn close notify: %v", err)
				s.isConnected.Store(false)
			}
		}
		time.Sleep(time.Minute)
	}
}
