package kafka

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/emberfarkas/pkg/log"
	"github.com/emberfarkas/pkg/queue"
	"github.com/emberfarkas/pkg/tracing"
	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// TracingProducer 生产者
type TracingProducer struct {
	tracer *tracing.Tracer
	pub    *kafka.Writer
	topic  string
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
	tracer := tracing.NewTracer(trace.SpanKindProducer, tracing.WithPropagator(
		propagation.NewCompositeTextMapPropagator(tracing.Metadata{}, propagation.Baggage{}, tracing.TraceContext{}),
	))
	pub := kafka.NewWriter(config)
	tracingPub := &TracingProducer{
		tracer: tracer,
		pub:    pub,
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
	ctx, span := p.tracer.Start(ctx, operation, &KafkaMessageTextMapCarrier{msg: msg})
	span.SetAttributes(
		attribute.String("kafka.topic", p.topic),
		attribute.String("kafka.key", string(msg.Key)),
	)
	err := p.pub.WriteMessages(ctx, msg)
	p.tracer.End(ctx, span, msg, err)
	err = WrapError(err)
	return err
}

func (p *TracingProducer) Close() error {
	return p.pub.Close()
}
