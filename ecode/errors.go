package ecode

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
)

// WrapError 转换错误
func WrapError(err error) error {
	if err == nil {
		return nil
	}
	// 同类型
	if se := new(errors.Error); errors.As(err, &se) {
		return err
	}
	return Unknown(err.Error())
}

// NewMetadataf NewMetadataf(code fmt.Sprintf(format, a...))
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
