package ecode

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
)

func Unknown(format string, a ...interface{}) error {
	return errors.InternalServer("Unknown", fmt.Sprintf(format, a...))
}

func InvalidParam(format string, a ...interface{}) error {
	return errors.InternalServer("InvalidParam", fmt.Sprintf(format, a...))
}

// main ecode interval is [0,990000]
func SignCheckErr(format string, a ...interface{}) error {
	return errors.InternalServer("10001", fmt.Sprintf(format, a...))
}

// ReadTimeout 系统错误
func ReadTimeout(format string, a ...interface{}) error {
	return errors.InternalServer("10100", "读取超时")
}

// admin jwt
// ErrMissingSecretKey indicates Secret key is required
func ErrMissingSecretKey(domain, reason, format string, a ...interface{}) error {
	return errors.InternalServer("11101", "secret key is required")
}

func ErrForbidden(domain, reason, format string, a ...interface{}) error {
	return errors.InternalServer("11102", "禁止") // ErrForbidden when HTTP status 403 is give
}

func ErrUpdateSysUser(domain, reason, format string, a ...interface{}) error {
	return errors.InternalServer("11103", "") // ErrForbidden when HTTP status 403 is give
}

func ErrAesNewCipher(domain, reason, format string, a ...interface{}) error {
	return errors.InternalServer("11104", "") // ErrForbidden when HTTP status 403 is give
}

func AdminTableEmpty(domain, reason, format string, a ...interface{}) error {
	return errors.InternalServer("11105", "") // ErrForbidden when HTTP status 403 is give
}

func ErrGetSysUser(domain, reason, format string, a ...interface{}) error {
	return errors.InternalServer("11107", "") // ErrForbidden when HTTP status 403 is give
}

func ErrUsername(domain, reason, format string, a ...interface{}) error {
	return errors.InternalServer("11108", "") // ErrForbidden when HTTP status 403 is give
}

func ErrNoPrivKeyFile(domain, reason, format string, a ...interface{}) error {
	return errors.InternalServer("11109", "") // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidPrivKey(domain, reason, format string, a ...interface{}) error {
	return errors.InternalServer("11110", "") // ErrForbidden when HTTP status 403 is give
}

func ErrUnsupportedResponseType(domain, reason, format string, a ...interface{}) error {
	return errors.InternalServer("11117", "") // ErrForbidden when HTTP status 403 is give
}

func ErrUnauthorizedClient(domain, reason, format string, a ...interface{}) error {
	return errors.InternalServer("11118", fmt.Sprintf(format, a...)) // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidScope(domain, reason, format string, a ...interface{}) error {
	return errors.InternalServer("11119", fmt.Sprintf(format, a...)) // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidUserId(domain, reason, format string, a ...interface{}) error {
	return errors.InternalServer("11120", fmt.Sprintf(format, a...)) // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidRequest(domain, reason, format string, a ...interface{}) error {
	return errors.InternalServer("11121", "") // ErrForbidden when HTTP status 403 is give
}

func ErrUnsupportedGrantType(domain, reason, format string, a ...interface{}) error {
	return errors.InternalServer("11122", "") // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidUsername(domain, reason, format string, a ...interface{}) error {
	return errors.InternalServer("11123", "") // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidUid(domain, reason, format string, a ...interface{}) error {
	return errors.InternalServer("11124", "") // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidAuthorizeCode(domain, reason, format string, a ...interface{}) error {
	return errors.InternalServer("11125", "") // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidClient(domain, reason, format string, a ...interface{}) error {
	return errors.InternalServer("11126", "") // ErrForbidden when HTTP status 403 is give
}

func ErrUnsupportedErrType(domain, reason, format string, a ...interface{}) error {
	return errors.InternalServer("11127", "") // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidGrant(domain, reason, format string, a ...interface{}) error {
	return errors.InternalServer("11128", "") // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidAccessToken(domain, reason, format string, a ...interface{}) error {
	return errors.InternalServer("11129", "") // ErrForbidden when HTTP status 403 is give
}

func UserNotExist(domain, reason, format string, a ...interface{}) error {
	return errors.InternalServer("11130", "") // ErrForbidden when HTTP status 403 is give
}

func ErrSignService(format string, a ...interface{}) error {
	return errors.InternalServer("11131", fmt.Sprintf(format, a...))
}

func IsErrSignService(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "11131" && se.Code == 500
}

func ErrUnimplemented(format string, a ...interface{}) error {
	return errors.InternalServer("Unimplemented", fmt.Sprintf(format, a...))
}

func IsUnimplemented(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "Unimplemented" && se.Code == 500
}

func ErrInvalidTransport(format string, a ...interface{}) error {
	return errors.InternalServer("ErrInvalidTransport", fmt.Sprintf(format, a...))
}

func IsErrInvalidTransport(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "ErrInvalidTransport" && se.Code == 500
}

func ErrInvalidPayload(format string, a ...interface{}) error {
	return errors.InternalServer("ErrInvalidPayload", fmt.Sprintf(format, a...))
}

func IsErrInvalidPayload(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "ErrInvalidPayload" && se.Code == 500
}

func ErrSocksProxy(format string, a ...interface{}) error {
	return errors.InternalServer("ErrSocksProxy", fmt.Sprintf(format, a...))
}
