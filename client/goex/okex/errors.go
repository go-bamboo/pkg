package okex

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
)

func ErrNotSupportType(format string, a ...interface{}) error {
	return errors.InternalServer("NotSupportType", fmt.Sprintf(format, a...))
}

func IsErrNotSupportType(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "NotSupportType" && se.Code == 500
}
