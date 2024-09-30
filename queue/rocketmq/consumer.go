package rocketmq

import (
	"context"
	"fmt"

	v2 "github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/go-bamboo/pkg/log"
	otelext "github.com/go-bamboo/pkg/otel"
	"github.com/go-bamboo/pkg/queue"
	"github.com/go-bamboo/pkg/rescue"
	"github.com/go-kratos/kratos/v2/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

func init() {
	queue.RegisterConsumer("redis", NewQueue)
}

type rocketQueue struct {
	c          *queue.Conf
	handler    queue.ConsumeHandler
	sub        v2.PushConsumer
	tracer     trace.Tracer
	propagator propagation.TextMapPropagator
	subCounter metric.Int64Counter // 发送次数
}

func MustNewQueue(c *queue.Conf, handler queue.ConsumeHandler) queue.MessageQueue {
	q, err := NewQueue(c, handler)
	if err != nil {
		log.Fatal(err)
	}
	return q
}

func NewQueue(c *queue.Conf, handler queue.ConsumeHandler) (queue.MessageQueue, error) {
	cc, err := newKafkaQueue(c, handler)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return cc, nil
}

func newKafkaQueue(config *queue.Conf, handler queue.ConsumeHandler) (k *rocketQueue, err error) {
	model := consumer.Clustering
	if config.Broadcast {
		model = consumer.BroadCasting
	}
	namesrvAdd, err := primitive.NewNamesrvAddr(config.Brokers...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	cs, err := v2.NewPushConsumer(
		consumer.WithGroupName(config.Group),
		consumer.WithNameServer(namesrvAdd),
		consumer.WithCredentials(primitive.Credentials{
			//AccessKey: config.AccessKey,
			//SecretKey: config.SecretKey,
		}),
		consumer.WithConsumerModel(model),
		//consumer.WithNamespace(config.Namespace),
		consumer.WithConsumeMessageBatchMaxSize(1), //
	)
	if err != nil {
		return nil, errors.FromError(err)
	}
	k = &rocketQueue{
		c:          config,
		handler:    handler,
		sub:        cs,
		tracer:     otel.Tracer("rocketmq"),
		propagator: propagation.NewCompositeTextMapPropagator(otelext.Metadata{}, propagation.Baggage{}, otelext.TraceContext{}),
	}
	return
}

func (c *rocketQueue) Name() string {
	return "rocketmq"
}

func (c *rocketQueue) Start(context.Context) error {
	log.Infof("start consumer topic:%v", c.c.Topic)
	if err := c.consumeGroupTopic(c.c.Topic, c.c.Expression); err != nil {
		return err
	}
	if err := c.sub.Start(); err != nil {
		log.Error(err)
		return err
	}
	log.Infof("start rocket consumer.")
	return nil
}

func (c *rocketQueue) Stop(context.Context) error {
	if err := c.sub.Shutdown(); err != nil {
		return err
	}
	log.Infof("stop rocket consumer.")
	return nil
}

func (c *rocketQueue) consumeGroupTopic(topic, expression string) error {
	selector := consumer.MessageSelector{Type: consumer.TAG, Expression: expression}
	err := c.sub.Subscribe(topic, selector, c.handleMsg)
	if err != nil {
		return err
	}
	return nil
}

func (c *rocketQueue) handleMsg(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	defer rescue.Recover()
	for _, msg := range msgs {
		cctx, span := c.tracer.Start(ctx, "sub:"+msg.Topic, trace.WithSpanKind(trace.SpanKindConsumer))
		c.propagator.Inject(cctx, &MessageExtTextMapCarrier{msg: msg})
		span.SetAttributes(
			attribute.String("kafka.topic", msg.Topic),
			attribute.String("kafka.key", string(msg.GetKeys())),
		)
		if err := c.handler.Consume(ctx, msg.Topic, []byte(msg.GetKeys()), msg.Body); err != nil {
			span.RecordError(err)
			se := errors.FromError(err)
			log.Errorw(fmt.Sprintf("%+v", err), "code", se.Code, "reason", se.Reason)
			return consumer.ConsumeRetryLater, err
		}
	}
	return consumer.ConsumeSuccess, nil
}
