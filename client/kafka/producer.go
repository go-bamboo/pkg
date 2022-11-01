package kafka

import (
	"context"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/queue"
	"github.com/go-bamboo/pkg/tracing"
	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// Producer 生产者
type Producer struct {
	pub        *kafka.Writer
	topic      string
	tracer     trace.Tracer
	propagator propagation.TextMapPropagator
}

func MustNewProducer(c *Conf) queue.Pusher {
	pub, err := NewProducer(c)
	if err != nil {
		log.Fatal(err)
	}
	return pub
}

func NewProducer(c *Conf) (*Producer, error) {
	pub := kafka.Writer{
		Addr: kafka.TCP("localhost:9092", "localhost:9093", "localhost:9094"),
	}
	tracingPub := &Producer{
		tracer:     otel.Tracer("kafka"),
		propagator: propagation.NewCompositeTextMapPropagator(tracing.Metadata{}, propagation.Baggage{}, tracing.TraceContext{}),
		pub:        &pub,
	}
	return tracingPub, nil
}

func (p *Producer) Name() string {
	return "kafka"
}

func (p *Producer) Push(ctx context.Context, topic string, key, value []byte) error {
	msg := kafka.Message{
		Key:   key,
		Value: value,
	}
	operation := "pub:" + topic
	ctx, span := p.tracer.Start(ctx, operation, trace.WithSpanKind(trace.SpanKindProducer))
	p.propagator.Inject(ctx, &KafkaMessageTextMapCarrier{msg: msg})
	span.SetAttributes(
		attribute.String("kafka.topic", topic),
		attribute.String("kafka.key", string(msg.Key)),
	)
	if err := p.pub.WriteMessages(ctx, msg); err != nil {
		span.RecordError(err)
		err = WrapError(err)
		return err
	}
	return nil
}

func (p *Producer) Close() error {
	return p.pub.Close()
}
