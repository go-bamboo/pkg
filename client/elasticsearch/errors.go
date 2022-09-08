package elasticsearch

import (
	"github.com/emberfarkas/pkg/ecode"
	"github.com/go-kratos/kratos/v2/errors"
)

func WrapEsError(err error) error {
	if err == nil {
		return nil
	}
	// 同类型
	if se := new(errors.Error); errors.As(err, &se) {
		return err
	}
	return ecode.Unknown(err.Error())
}

func ErrEs(msg string) error {
	return errors.InternalServer("ErrEs", msg)
}

func IsErrEs(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "ErrEs" && se.Code == 500
}
