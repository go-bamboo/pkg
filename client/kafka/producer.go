package kafka

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/emberfarkas/pkg/log"
	"github.com/emberfarkas/pkg/queue"
	"github.com/emberfarkas/pkg/tracing"
	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// TracingProducer 生产者
type TracingProducer struct {
	pub        *kafka.Writer
	topic      string
	tracer     trace.Tracer
	propagator propagation.TextMapPropagator
}

func MustNewTracingProducer(c *Conf) queue.Pusher {
	pub, err := NewTracingProducer(c)
	if err != nil {
		log.Fatal(err)
	}
	return pub
}

func NewTracingProducer(c *Conf) (*TracingProducer, error) {
	dialer := &Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
		TLS:       &tls.Config{},
	}
	config := kafka.WriterConfig{
		Brokers:  c.Brokers,
		Dialer:   (*kafka.Dialer)(dialer),
		Balancer: &kafka.LeastBytes{},
	}
	pub := kafka.NewWriter(config)
	tracingPub := &TracingProducer{
		tracer:     otel.Tracer("kafka"),
		propagator: propagation.NewCompositeTextMapPropagator(tracing.Metadata{}, propagation.Baggage{}, tracing.TraceContext{}),
		pub:        pub,
	}
	return tracingPub, nil
}

func (p *TracingProducer) Name() string {
	return ""
}

func (p *TracingProducer) Push(ctx context.Context, key, value []byte) error {
	msg := kafka.Message{
		Key:   key,
		Value: value,
	}
	operation := "pub:" + msg.Topic
	ctx, span := p.tracer.Start(ctx, operation, trace.WithSpanKind(trace.SpanKindProducer))
	p.propagator.Inject(ctx, &KafkaMessageTextMapCarrier{msg: msg})
	span.SetAttributes(
		attribute.String("kafka.topic", p.topic),
		attribute.String("kafka.key", string(msg.Key)),
	)
	err := p.pub.WriteMessages(ctx, msg)
	if err != nil {
		span.RecordError(err)
		err = WrapError(err)
	}
	return err
}

func (p *TracingProducer) Close() error {
	return p.pub.Close()
}
