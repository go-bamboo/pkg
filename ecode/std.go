package ecode

import (
	"github.com/go-kratos/kratos/v2/errors"
)

func WrapStdError(err error) error {
	if err == nil {
		return nil
	}
	return errors.InternalServer("StdErr", err.Error())
}

func IsStdErr(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "StdErr" && se.Code == 500
}
