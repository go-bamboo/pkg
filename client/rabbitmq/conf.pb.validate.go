// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: client/rabbitmq/conf.proto

package rabbitmq

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on RabbitConf with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *RabbitConf) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Username

	// no validation rules for Password

	// no validation rules for Host

	// no validation rules for Port

	// no validation rules for VHost

	// no validation rules for Address

	return nil
}

// RabbitConfValidationError is the validation error returned by
// RabbitConf.Validate if the designated constraints aren't met.
type RabbitConfValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RabbitConfValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RabbitConfValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RabbitConfValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RabbitConfValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RabbitConfValidationError) ErrorName() string { return "RabbitConfValidationError" }

// Error satisfies the builtin error interface
func (e RabbitConfValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRabbitConf.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RabbitConfValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RabbitConfValidationError{}

// Validate checks the field values on ConsumerConf with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ConsumerConf) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Consumer

	return nil
}

// ConsumerConfValidationError is the validation error returned by
// ConsumerConf.Validate if the designated constraints aren't met.
type ConsumerConfValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConsumerConfValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConsumerConfValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConsumerConfValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConsumerConfValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConsumerConfValidationError) ErrorName() string { return "ConsumerConfValidationError" }

// Error satisfies the builtin error interface
func (e ConsumerConfValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConsumerConf.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConsumerConfValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConsumerConfValidationError{}

// Validate checks the field values on ListenerConf with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListenerConf) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRabbit()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerConfValidationError{
				field:  "Rabbit",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetQueues() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListenerConfValidationError{
					field:  fmt.Sprintf("Queues[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListenerConfValidationError is the validation error returned by
// ListenerConf.Validate if the designated constraints aren't met.
type ListenerConfValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListenerConfValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListenerConfValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListenerConfValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListenerConfValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListenerConfValidationError) ErrorName() string { return "ListenerConfValidationError" }

// Error satisfies the builtin error interface
func (e ListenerConfValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListenerConf.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListenerConfValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListenerConfValidationError{}

// Validate checks the field values on ProducerConf with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ProducerConf) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRabbit()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProducerConfValidationError{
				field:  "Rabbit",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ContentType

	return nil
}

// ProducerConfValidationError is the validation error returned by
// ProducerConf.Validate if the designated constraints aren't met.
type ProducerConfValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProducerConfValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProducerConfValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProducerConfValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProducerConfValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProducerConfValidationError) ErrorName() string { return "ProducerConfValidationError" }

// Error satisfies the builtin error interface
func (e ProducerConfValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProducerConf.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProducerConfValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProducerConfValidationError{}

// Validate checks the field values on ExchangeConf with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ExchangeConf) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Exchange

	// no validation rules for Key

	return nil
}

// ExchangeConfValidationError is the validation error returned by
// ExchangeConf.Validate if the designated constraints aren't met.
type ExchangeConfValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExchangeConfValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExchangeConfValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExchangeConfValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExchangeConfValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExchangeConfValidationError) ErrorName() string { return "ExchangeConfValidationError" }

// Error satisfies the builtin error interface
func (e ExchangeConfValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExchangeConf.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExchangeConfValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExchangeConfValidationError{}

// Validate checks the field values on QueueConf with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *QueueConf) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	return nil
}

// QueueConfValidationError is the validation error returned by
// QueueConf.Validate if the designated constraints aren't met.
type QueueConfValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueueConfValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueueConfValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueueConfValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueueConfValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueueConfValidationError) ErrorName() string { return "QueueConfValidationError" }

// Error satisfies the builtin error interface
func (e QueueConfValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQueueConf.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueueConfValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueueConfValidationError{}
