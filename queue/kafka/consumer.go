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

// Name is the name registered for kafka
const Name = "kafka"

func init() {
	queue.RegisterConsumer(Name, NewConsumer)
	queue.RegisterPusher(Name, NewProducer)
}

type Consumer struct {
	c       *queue.Conf
	handler map[string]queue.ConsumeHandle
	wg      sync.WaitGroup
	ctx     context.Context
	cf      context.CancelFunc
}

func MustNewQueue(c *queue.Conf) (queue.MessageQueue, error) {
	q, err := NewConsumer(c)
	if err != nil {
		log.Fatal(err)
	}
	return q, nil
}

func NewConsumer(c *queue.Conf) (queue.MessageQueue, error) {
	ctx, cf := context.WithCancel(context.Background())
	tracingSub := &Consumer{
		c:       c,
		handler: map[string]queue.ConsumeHandle{},
		ctx:     ctx,
		cf:      cf,
	}
	log.Infof("start kafka consumer")
	return tracingSub, nil
}

func (c *Consumer) Name() string {
	return "kafka"
}

func (c *Consumer) Subscribe(topic string, handler queue.ConsumeHandle, opts ...queue.SubscribeOption) (queue.Subscriber, error) {
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
	config := kafka.ReaderConfig{
		Brokers: c.c.Brokers,
		GroupID: c.c.Group,
		Topic:   topic,
		Dialer:  &dialer,
	}
	sub := kafka.NewReader(config)
	s := Subscriber{
		topic:      topic,
		sub:        sub,
		tracer:     otel.Tracer("kafka"),
		propagator: propagation.NewCompositeTextMapPropagator(otelext.Metadata{}, propagation.Baggage{}, otelext.TraceContext{}),
		handler:    handler,
	}
	c.wg.Add(1)
	go c.consumGroupTopic(c.ctx, &s)
	return &s, nil
}

func (c *Consumer) Close() error {
	c.cf()
	c.wg.Wait()
	log.Info("stop kafka consumer. topic")
	return nil
}

func (c *Consumer) poll(ctx context.Context, sub *Subscriber, timeoutMs int) (cctx context.Context, span trace.Span, msg kafka.Message, err error) {
	msg, err = sub.sub.ReadMessage(ctx)
	if err != nil {
		return
	}
	cctx, span = sub.tracer.Start(ctx, "sub:"+msg.Topic, trace.WithSpanKind(trace.SpanKindConsumer))
	sub.propagator.Inject(ctx, &KafkaMessageTextMapCarrier{msg: msg})
	span.SetAttributes(
		attribute.String("kafka.topic", msg.Topic),
		attribute.String("kafka.key", string(msg.Key)),
	)
	return
}

func (c *Consumer) consumGroupTopic(ctx context.Context, sub *Subscriber) {
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
			cCtx, span, msg, err := c.poll(cCtx, sub, 100)
			if err != nil {
				log.Errorf("err: %v", err)
				cf()
				continue
			}
			handler := c.handler[msg.Topic]
			if err := handler(cCtx, msg.Topic, msg.Key, msg.Value); err != nil {
				// 直接放弃的消息
				se := errors.FromError(err)
				log.Errorw(fmt.Sprintf("%+v", err), "code", se.Code, "reason", se.Reason, "topic", msg.Topic, "partition", msg.Partition, "offset", msg.Offset)
			}
			// 确认消费
			if err := sub.sub.CommitMessages(ctx, msg); err != nil {
				span.RecordError(err)
			}
			cf()
		}
	}
}
