package ecode

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
)

// NewMetadataf(code fmt.Sprintf(format, a...))
func NewMetadataf(code int, reason string, md map[string]string, format string, a ...interface{}) error {
	se := errors.New(code, reason, fmt.Sprintf(format, a...))
	se = se.WithMetadata(md)
	return se
}

// IsTimeout checks whether the given error is a timeout.
func IsTimeout(err error) bool {
	timeoutErr, ok := err.(interface {
		Timeout() bool
	})
	return ok && timeoutErr.Timeout()
}
