package kafka

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/go-bamboo/pkg/log"
	otelext "github.com/go-bamboo/pkg/otel"
	"github.com/go-bamboo/pkg/queue"
	"github.com/go-bamboo/pkg/rescue"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type (
	ConsumeHandle func(ctx context.Context, topic string, key, message []byte) error

	ConsumeHandler interface {
		Consume(ctx context.Context, topic string, key, message []byte) error
	}
	Consumer struct {
		c       *Conf
		handler ConsumeHandler

		sub        *kafka.Reader
		tracer     trace.Tracer
		propagator propagation.TextMapPropagator

		wg  sync.WaitGroup
		ctx context.Context
		cf  context.CancelFunc
	}

	Consumers struct {
		queues []*Consumer
	}
)

func MustNewQueue(c *Conf, handler ConsumeHandler) (queue.MessageQueue, error) {
	q, err := NewQueue(c, handler)
	if err != nil {
		log.Fatal(err)
	}
	return q, nil
}

func NewQueue(c *Conf, handler ConsumeHandler) (*Consumers, error) {
	q := Consumers{}

	cc, err := NewConsumer(c, handler)
	if err != nil {
		return nil, err
	}
	q.queues = append(q.queues, cc)
	return &q, nil
}

func (q Consumers) Name() string {
	return "kafka"
}

func (q Consumers) Start(ctx context.Context) error {
	for _, queue := range q.queues {
		queue.Start(ctx)
	}
	return nil
}

func (q Consumers) Stop(ctx context.Context) error {
	for _, queue := range q.queues {
		queue.Stop(ctx)
	}
	return nil
}

func NewConsumer(c *Conf, handler ConsumeHandler) (*Consumer, error) {
	// Load client cert
	//cert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	//if err != nil {
	//	return &tlsConfig, err
	//}
	//tlsConfig.Certificates = []tls.Certificate{cert}
	// load ca
	bytes, err := os.ReadFile("/mnt/d/Documents/GitHub/goblockchian/blockchian/server/cert/phy_ca.crt")
	if err != nil {
		return nil, err
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(bytes)
	dialer := kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
		SASLMechanism: plain.Mechanism{
			Username: "kafkauser",
			Password: "Kafkauser-music-bee",
		},
		TLS: &tls.Config{
			RootCAs:            pool,
			InsecureSkipVerify: true,
		},
	}
	ctx, cf := context.WithCancel(context.Background())
	config := kafka.ReaderConfig{
		Brokers: c.Brokers,
		GroupID: c.Group,
		Topic:   c.Topic,
		Dialer:  &dialer,
	}
	sub := kafka.NewReader(config)
	tracingSub := &Consumer{
		c:       c,
		handler: handler,

		sub:        sub,
		tracer:     otel.Tracer("kafka"),
		propagator: propagation.NewCompositeTextMapPropagator(otelext.Metadata{}, propagation.Baggage{}, otelext.TraceContext{}),

		ctx: ctx,
		cf:  cf,
	}
	return tracingSub, nil
}

func (c *Consumer) poll(ctx context.Context, timeoutMs int) (cctx context.Context, span trace.Span, msg kafka.Message, err error) {
	msg, err = c.sub.ReadMessage(ctx)
	if err != nil {
		return
	}
	cctx, span = c.tracer.Start(ctx, "sub:"+msg.Topic, trace.WithSpanKind(trace.SpanKindConsumer))
	c.propagator.Inject(ctx, &KafkaMessageTextMapCarrier{msg: msg})
	span.SetAttributes(
		attribute.String("kafka.topic", msg.Topic),
		attribute.String("kafka.key", string(msg.Key)),
	)
	return
}

func (c *Consumer) Start(context.Context) error {
	c.wg.Add(1)
	go c.consumGroupTopic(c.ctx)
	log.Infof("start kafka consumer, topic[%s]", c.c.Topic)
	return nil
}

func (c *Consumer) Stop(context.Context) error {
	c.cf()
	c.wg.Wait()
	c.sub.Close()
	log.Info("stop kafka consumer. topic[%s]", c.c.Topic)
	return nil
}

func (c *Consumer) consumGroupTopic(ctx context.Context) {
	defer rescue.Recover(func() {
		c.wg.Done()
		log.Warnf("kafka consumGroupTopic done")
	})
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// ms
			cCtx, cf := context.WithTimeout(context.TODO(), 60*time.Second)
			cCtx, span, msg, err := c.poll(cCtx, 100)
			if err != nil {
				log.Errorf("err: %v", err)
				cf()
				continue
			}
			if err := c.handler.Consume(cCtx, c.c.Topic, msg.Key, msg.Value); err != nil {
				// 直接放弃的消息
				se := errors.FromError(err)
				log.Errorw(fmt.Sprintf("%+v", err), "code", se.Code, "reason", se.Reason, "topic", msg.Topic, "partition", msg.Partition, "offset", msg.Offset)
			}
			// 确认消费
			if err := c.sub.CommitMessages(ctx, msg); err != nil {
				span.RecordError(err)
			}
			cf()
		}
	}
}
