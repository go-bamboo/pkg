package rabbitmq

import (
	"context"
	"sync"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/queue"
	"github.com/go-bamboo/pkg/rescue"
	amqp "github.com/rabbitmq/amqp091-go"
)

// Name is the name registered for kafka
const Name = "rabbitmq"

func init() {
	queue.RegisterConsumer(Name, NewListener)
	queue.RegisterPusher(Name, NewSender)
}

type RabbitListener struct {
	c       *queue.Conf
	topic   chan string
	handler map[string]queue.ConsumeHandle
	ctx     context.Context
	cf      context.CancelFunc
	wg      sync.WaitGroup
}

func MustNewListener(c *queue.Conf) queue.MessageQueue {
	listener, err := NewListener(c)
	if err != nil {
		log.Fatal(err)
	}
	return listener
}

func NewListener(c *queue.Conf) (consumer queue.MessageQueue, err error) {
	ctx, cf := context.WithCancel(context.Background())
	listener := &RabbitListener{
		c:       c,
		topic:   make(chan string),
		handler: make(map[string]queue.ConsumeHandle),
		ctx:     ctx,
		cf:      cf,
	}
	listener.wg.Add(1)
	go listener.reconnect()
	log.Infof("[rabbitmq][listener] start consume queue.")
	return listener, nil
}

func (s *RabbitListener) Name() string {
	return "rabbitListener"
}

func (s *RabbitListener) Subscribe(topic string, handler queue.ConsumeHandle, opts ...queue.SubscribeOption) (queue.Subscriber, error) {
	s.handler[topic] = handler
	s.topic <- topic
	return nil, nil
}

func (s *RabbitListener) Close() error {
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
	handler := s.handler[topic]
	return handler(c, topic, []byte(msg.RoutingKey), d)
}
