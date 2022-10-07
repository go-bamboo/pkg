package kafka

import (
	"context"
	"crypto/tls"
	"fmt"
	"sync"
	"time"

	"github.com/emberfarkas/pkg/log"
	"github.com/emberfarkas/pkg/queue"
	"github.com/emberfarkas/pkg/rescue"
	"github.com/emberfarkas/pkg/stat/prom"
	"github.com/emberfarkas/pkg/tracing"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type (
	TracingConsumer struct {
		c       Conf
		handler queue.ConsumeHandler

		sub     *kafka.Reader
		tracer  *tracing.Tracer
		metrics *prom.Prom

		wg  sync.WaitGroup
		ctx context.Context
		cf  context.CancelFunc
	}

	TracingConsumers struct {
		queues []*TracingConsumer
	}
)

func MustNewQueue(c *Conf, handler queue.ConsumeHandler) (queue.MessageQueue, error) {
	q, err := NewQueue(c, handler)
	if err != nil {
		log.Fatal(err)
	}
	return q, nil
}

func NewQueue(c *Conf, handler queue.ConsumeHandler) (*TracingConsumers, error) {
	q := TracingConsumers{}
	dialer := &Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
		TLS:       &tls.Config{},
	}
	cc, err := NewTracingConsumer(*c, dialer, handler)
	if err != nil {
		return nil, err
	}
	q.queues = append(q.queues, cc)
	return &q, nil
}

func (q TracingConsumers) Name() string {
	return ""
}

func (q TracingConsumers) Start(ctx context.Context) error {
	for _, queue := range q.queues {
		queue.Start(ctx)
	}
	return nil
}

func (q TracingConsumers) Stop(ctx context.Context) error {
	for _, queue := range q.queues {
		queue.Stop(ctx)
	}
	return nil
}

func NewTracingConsumer(c Conf, dialer *Dialer, handler queue.ConsumeHandler) (*TracingConsumer, error) {
	ctx, cf := context.WithCancel(context.Background())
	config := kafka.ReaderConfig{
		Brokers: c.Brokers,
		Topic:   c.Topic,
		Dialer:  (*kafka.Dialer)(dialer),
	}
	tracer := tracing.NewTracer(trace.SpanKindConsumer, tracing.WithPropagator(
		propagation.NewCompositeTextMapPropagator(tracing.Metadata{}, propagation.Baggage{}, tracing.TraceContext{}),
	))
	sub := kafka.NewReader(config)
	tracingSub := &TracingConsumer{
		c:       c,
		sub:     sub,
		handler: handler,

		tracer: tracer,

		ctx: ctx,
		cf:  cf,
	}
	return tracingSub, nil
}

func (c *TracingConsumer) poll(ctx context.Context, timeoutMs int) (cctx context.Context, out kafka.Message, err error) {
	msg, err := c.sub.ReadMessage(ctx)
	if err != nil {
		return
	}
	ctx, span := c.tracer.Start(ctx, "sub:"+msg.Topic, &KafkaMessageTextMapCarrier{msg: msg})
	span.SetAttributes(
		attribute.String("kafka.topic", msg.Topic),
		attribute.String("kafka.key", string(msg.Key)),
	)
	cctx = context.WithValue(ctx, TracingConsumer{}, span)
	out = kafka.Message(msg)
	return
}

func (c *TracingConsumer) commitMessage(ctx context.Context, m kafka.Message) error {
	err := c.sub.CommitMessages(ctx, kafka.Message(m))
	span, ok := ctx.Value(TracingConsumer{}).(trace.Span)
	if ok {
		c.tracer.End(ctx, span, m, nil)
	}
	return err
}

func (s *TracingConsumer) Start(context.Context) error {
	s.wg.Add(1)
	go s.consumGroupTopic(s.ctx)
	log.Infof("start kafka consumer, topic[%s]", s.c.Topic)
	return nil
}

func (s *TracingConsumer) Stop(context.Context) error {
	s.cf()
	s.wg.Wait()
	s.sub.Close()
	log.Info("stop kafka consumer. topic[%s]", s.c.Topic)
	return nil
}

func (s *TracingConsumer) consumGroupTopic(ctx context.Context) {
	defer rescue.Recover(func() {
		s.wg.Done()
		log.Warnf("kafka consumGroupTopic done")
	})
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// ms
			cCtx, cf := context.WithTimeout(context.TODO(), 60*time.Second)
			cCtx, msg, err := s.poll(ctx, 100)
			if err != nil {
				log.Errorf("err: %v", err)
				cf()
				continue
			}
			if err := s.handler.Consume(cCtx, s.c.Topic, msg.Key, msg.Value); err != nil {
				// 直接放弃的消息
				se := errors.FromError(err)
				log.Errorw(fmt.Sprintf("%+v", err), "code", se.Code, "reason", se.Reason, "topic", msg.Topic, "partition", msg.Partition, "offset", msg.Offset)
			}
			// 确认消费
			if err := s.commitMessage(cCtx, msg); err != nil {
				log.Errorf("err: %v", err)
			}
			cf()
		}
	}
}
