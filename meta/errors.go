package meta

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
)

func ErrNotExistMd(format string, a ...interface{}) error {
	return errors.InternalServer("ErrNotExistMd", fmt.Sprintf(format, a...))
}

func IsErrNotExistMd(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "ErrNotExistMd" && se.Code == 500
}

func ErrNotExistDp(format string, a ...interface{}) error {
	return errors.InternalServer("ErrNotExistDp", fmt.Sprintf(format, a...))
}

func IsErrNotExistDp(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "ErrNotExistDp" && se.Code == 500
}
