package rocketmq

import (
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"go.opentelemetry.io/otel/propagation"
)

type MessageTextMapCarrier struct {
	msg *primitive.Message
}

var _ propagation.TextMapCarrier = &MessageTextMapCarrier{}

// Get returns the value associated with the passed key.
func (carrier *MessageTextMapCarrier) Get(key string) string {
	//for i := 0; i < len(carrier.msg.); i++ {
	//	header := carrier.msg.Headers[i]
	//	if strings.Compare("md-"+key, header.Key) == 0 {
	//		return string(header.Value)
	//	}
	//}
	return ""
}

// Set stores the key-value pair.
func (carrier *MessageTextMapCarrier) Set(key string, value string) {
	//carrier.msg.WithProperty(key, value)
	//carrier.msg.Headers = append(carrier.msg.Headers, kafka.Header{
	//	Key:   "md-" + key,
	//	Value: []byte(value),
	//})
}

// Keys lists the keys stored in this carrier.
func (carrier *MessageTextMapCarrier) Keys() []string {
	return nil
}

type MessageExtTextMapCarrier struct {
	msg *primitive.MessageExt
}

// Get returns the value associated with the passed key.
func (carrier *MessageExtTextMapCarrier) Get(key string) string {
	//for i := 0; i < len(carrier.msg.); i++ {
	//	header := carrier.msg.Headers[i]
	//	if strings.Compare("md-"+key, header.Key) == 0 {
	//		return string(header.Value)
	//	}
	//}
	return ""
}

// Set stores the key-value pair.
func (carrier *MessageExtTextMapCarrier) Set(key string, value string) {
	//carrier.msg.WithProperty(key, value)
	//carrier.msg.Headers = append(carrier.msg.Headers, kafka.Header{
	//	Key:   "md-" + key,
	//	Value: []byte(value),
	//})
}

// Keys lists the keys stored in this carrier.
func (carrier *MessageExtTextMapCarrier) Keys() []string {
	return nil
}
