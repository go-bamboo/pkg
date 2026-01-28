package kafka

import (
	"github.com/go-bamboo/pkg/queue"
	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type Subscriber struct {
	topic      string
	sub        *kafka.Reader
	tracer     trace.Tracer
	propagator propagation.TextMapPropagator
	handler    queue.ConsumeHandle
}

// Options .
func (s *Subscriber) Options() queue.SubscribeOptions {
	return queue.SubscribeOptions{}
}

// Topic .
func (s *Subscriber) Topic() string {
	return s.topic
}

// Unsubscribe .
func (s *Subscriber) Unsubscribe(removeFromManager bool) error {
	return s.sub.Close()
}
