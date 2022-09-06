package kafka

import (
	"context"
	"crypto/tls"
	"time"

	"edu/pkg/log"
	"edu/pkg/tracing"

	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

//TracingProducer 生产者
type TracingProducer struct {
	tracer *tracing.Tracer
	pub    *kafka.Writer
	topic  string
}

func MustNewTracingProducer(c *Conf) *TracingProducer {
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

func (p *TracingProducer) Produce(ctx context.Context, key, value []byte) error {
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

func (p *TracingProducer) Close() {
	p.pub.Close()
}
