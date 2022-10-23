package confluent

import (
	"github.com/go-kratos/kratos/v2/errors"
)

func WrapError(err error) error {
	if err == nil {
		return nil
	}
	// 同类型
	if se := new(errors.Error); errors.As(err, &se) {
		return err
	}
	return errors.FromError(err)
}
