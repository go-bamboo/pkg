// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: registry/conf.proto

package registry

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on Etcd with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Etcd) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Etcd with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in EtcdMultiError, or nil if none found.
func (m *Etcd) ValidateAll() error {
	return m.validate(true)
}

func (m *Etcd) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Enable

	if all {
		switch v := interface{}(m.GetDialTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EtcdValidationError{
					field:  "DialTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EtcdValidationError{
					field:  "DialTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDialTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EtcdValidationError{
				field:  "DialTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return EtcdMultiError(errors)
	}

	return nil
}

// EtcdMultiError is an error wrapping multiple validation errors returned by
// Etcd.ValidateAll() if the designated constraints aren't met.
type EtcdMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EtcdMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EtcdMultiError) AllErrors() []error { return m }

// EtcdValidationError is the validation error returned by Etcd.Validate if the
// designated constraints aren't met.
type EtcdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EtcdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EtcdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EtcdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EtcdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EtcdValidationError) ErrorName() string { return "EtcdValidationError" }

// Error satisfies the builtin error interface
func (e EtcdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEtcd.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EtcdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EtcdValidationError{}

// Validate checks the field values on Consul with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Consul) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Consul with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ConsulMultiError, or nil if none found.
func (m *Consul) ValidateAll() error {
	return m.validate(true)
}

func (m *Consul) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Enable

	// no validation rules for Address

	if all {
		switch v := interface{}(m.GetDialTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConsulValidationError{
					field:  "DialTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConsulValidationError{
					field:  "DialTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDialTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConsulValidationError{
				field:  "DialTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ConsulMultiError(errors)
	}

	return nil
}

// ConsulMultiError is an error wrapping multiple validation errors returned by
// Consul.ValidateAll() if the designated constraints aren't met.
type ConsulMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConsulMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConsulMultiError) AllErrors() []error { return m }

// ConsulValidationError is the validation error returned by Consul.Validate if
// the designated constraints aren't met.
type ConsulValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConsulValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConsulValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConsulValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConsulValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConsulValidationError) ErrorName() string { return "ConsulValidationError" }

// Error satisfies the builtin error interface
func (e ConsulValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConsul.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConsulValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConsulValidationError{}

// Validate checks the field values on Kube with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Kube) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Kube with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in KubeMultiError, or nil if none found.
func (m *Kube) ValidateAll() error {
	return m.validate(true)
}

func (m *Kube) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Enable

	if len(errors) > 0 {
		return KubeMultiError(errors)
	}

	return nil
}

// KubeMultiError is an error wrapping multiple validation errors returned by
// Kube.ValidateAll() if the designated constraints aren't met.
type KubeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m KubeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m KubeMultiError) AllErrors() []error { return m }

// KubeValidationError is the validation error returned by Kube.Validate if the
// designated constraints aren't met.
type KubeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KubeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KubeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KubeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KubeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KubeValidationError) ErrorName() string { return "KubeValidationError" }

// Error satisfies the builtin error interface
func (e KubeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKube.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KubeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KubeValidationError{}

// Validate checks the field values on Nacos with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Nacos) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Nacos with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in NacosMultiError, or nil if none found.
func (m *Nacos) ValidateAll() error {
	return m.validate(true)
}

func (m *Nacos) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Enable

	// no validation rules for IpAddr

	// no validation rules for Port

	// no validation rules for Namespace

	if len(errors) > 0 {
		return NacosMultiError(errors)
	}

	return nil
}

// NacosMultiError is an error wrapping multiple validation errors returned by
// Nacos.ValidateAll() if the designated constraints aren't met.
type NacosMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NacosMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NacosMultiError) AllErrors() []error { return m }

// NacosValidationError is the validation error returned by Nacos.Validate if
// the designated constraints aren't met.
type NacosValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NacosValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NacosValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NacosValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NacosValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NacosValidationError) ErrorName() string { return "NacosValidationError" }

// Error satisfies the builtin error interface
func (e NacosValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNacos.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NacosValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NacosValidationError{}

// Validate checks the field values on Conf with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Conf) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Conf with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ConfMultiError, or nil if none found.
func (m *Conf) ValidateAll() error {
	return m.validate(true)
}

func (m *Conf) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetEtcd()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConfValidationError{
					field:  "Etcd",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConfValidationError{
					field:  "Etcd",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEtcd()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfValidationError{
				field:  "Etcd",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetConsul()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConfValidationError{
					field:  "Consul",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConfValidationError{
					field:  "Consul",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConsul()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfValidationError{
				field:  "Consul",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetKube()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConfValidationError{
					field:  "Kube",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConfValidationError{
					field:  "Kube",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetKube()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfValidationError{
				field:  "Kube",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetNacos()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConfValidationError{
					field:  "Nacos",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConfValidationError{
					field:  "Nacos",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNacos()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfValidationError{
				field:  "Nacos",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ConfMultiError(errors)
	}

	return nil
}

// ConfMultiError is an error wrapping multiple validation errors returned by
// Conf.ValidateAll() if the designated constraints aren't met.
type ConfMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConfMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConfMultiError) AllErrors() []error { return m }

// ConfValidationError is the validation error returned by Conf.Validate if the
// designated constraints aren't met.
type ConfValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfValidationError) ErrorName() string { return "ConfValidationError" }

// Error satisfies the builtin error interface
func (e ConfValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConf.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfValidationError{}
