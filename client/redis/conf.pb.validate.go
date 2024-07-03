// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: client/redis/conf.proto

package redis

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

// Validate checks the field values on Tls with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Tls) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Tls with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TlsMultiError, or nil if none found.
func (m *Tls) ValidateAll() error {
	return m.validate(true)
}

func (m *Tls) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for InsecureSkipVerify

	if len(errors) > 0 {
		return TlsMultiError(errors)
	}

	return nil
}

// TlsMultiError is an error wrapping multiple validation errors returned by
// Tls.ValidateAll() if the designated constraints aren't met.
type TlsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TlsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TlsMultiError) AllErrors() []error { return m }

// TlsValidationError is the validation error returned by Tls.Validate if the
// designated constraints aren't met.
type TlsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TlsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TlsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TlsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TlsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TlsValidationError) ErrorName() string { return "TlsValidationError" }

// Error satisfies the builtin error interface
func (e TlsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTls.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TlsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TlsValidationError{}

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
		switch v := interface{}(m.GetReadTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConfValidationError{
					field:  "ReadTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConfValidationError{
					field:  "ReadTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetReadTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfValidationError{
				field:  "ReadTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetWriteTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConfValidationError{
					field:  "WriteTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConfValidationError{
					field:  "WriteTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWriteTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfValidationError{
				field:  "WriteTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDialTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConfValidationError{
					field:  "DialTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConfValidationError{
					field:  "DialTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDialTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfValidationError{
				field:  "DialTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetExpireTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConfValidationError{
					field:  "ExpireTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConfValidationError{
					field:  "ExpireTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExpireTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfValidationError{
				field:  "ExpireTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Username

	// no validation rules for Password

	// no validation rules for Db

	if all {
		switch v := interface{}(m.GetTls()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConfValidationError{
					field:  "Tls",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConfValidationError{
					field:  "Tls",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTls()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfValidationError{
				field:  "Tls",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Debug

	// no validation rules for PoolSize

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
