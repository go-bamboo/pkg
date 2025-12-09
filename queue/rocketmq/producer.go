package rocketmq

import (
	"context"
	"fmt"

	v2 "github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/go-bamboo/pkg/log"
	otelext "github.com/go-bamboo/pkg/otel"
	"github.com/go-bamboo/pkg/queue"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// rocketProducer 生产者
type rocketProducer struct {
	producer   v2.Producer
	tracer     trace.Tracer
	propagator propagation.TextMapPropagator
	pubCounter metric.Int64Counter // 发送次数
	// topic      string
}

func MustNewPusher(c *queue.Conf) queue.Pusher {
	pub, err := NewPusher(c)
	if err != nil {
		log.Fatal(err)
	}
	return pub
}

func NewPusher(config *queue.Conf) (queue.Pusher, error) {
	pd, err := v2.NewProducer(
		producer.WithNameServer(config.Brokers),
		producer.WithGroupName(config.Group),
		producer.WithRetry(3),
		producer.WithCredentials(primitive.Credentials{
			//AccessKey: config.AccessKey,
			//SecretKey: config.SecretKey,
		}),
		producer.WithNamespace(config.Namespace),
	)
	if err != nil {
		return nil, fmt.Errorf("create new pd err:%s", err)
	}
	if err := pd.Start(); err != nil {
		return nil, err
	}
	tracingPub := &rocketProducer{
		producer:   pd,
		tracer:     otel.Tracer("roketmq"),
		propagator: propagation.NewCompositeTextMapPropagator(otelext.Metadata{}, propagation.Baggage{}, otelext.TraceContext{}),
		// topic:      c.Topic,
	}
	return tracingPub, nil
}

func (p *rocketProducer) Name() string {
	return "rocketmq"
}

func (p *rocketProducer) Push(ctx context.Context, topic string, key, value []byte) error {
	msg := primitive.NewMessage(topic, value)
	msg.WithTag("")
	msg.WithKeys([]string{string(key)})

	operation := "pub:" + topic
	ctx, span := p.tracer.Start(ctx, operation, trace.WithSpanKind(trace.SpanKindProducer))
	p.propagator.Inject(ctx, &MessageTextMapCarrier{msg: msg})
	span.SetAttributes(
		attribute.String("kafka.topic", topic),
		attribute.String("kafka.key", string(key)),
	)
	sendResult, err := p.producer.SendSync(ctx, msg)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Debugf("Delivered message to topic %s [%v] at offset %v", topic, sendResult.RegionID, sendResult.OffsetMsgID)
	return nil
}

func (p *rocketProducer) PushWithPartition(ctx context.Context, topic string, key, value []byte, partition int32) error {
	msg := primitive.NewMessage(topic, value)
	msg.WithTag("")
	msg.WithKeys([]string{string(key)})

	operation := "pub:" + topic
	ctx, span := p.tracer.Start(ctx, operation, trace.WithSpanKind(trace.SpanKindProducer))
	p.propagator.Inject(ctx, &MessageTextMapCarrier{msg: msg})
	span.SetAttributes(
		attribute.String("kafka.topic", topic),
		attribute.String("kafka.key", string(key)),
	)
	sendResult, err := p.producer.SendSync(ctx, msg)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Debugf("Delivered message to topic %s [%v] at offset %v", topic, sendResult.RegionID, sendResult.OffsetMsgID)
	return nil
}

func (p *rocketProducer) Close() error {
	return p.producer.Shutdown()
}
