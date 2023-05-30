// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package s3

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

// 为某个枚举单独设置错误码
func IsNotAllowExt(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NOT_ALLOW_EXT.String() && e.Code == 500
}

// 为某个枚举单独设置错误码
func ErrorNotAllowExt(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_NOT_ALLOW_EXT.String(), fmt.Sprintf(format, args...))
}

func IsConfigNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_CONFIG_NOT_FOUND.String() && e.Code == 500
}

func ErrorConfigNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_CONFIG_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}
