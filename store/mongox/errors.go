package mongox

import (
	"github.com/go-kratos/kratos/v2/errors"
)

// WrapGormError 转换错误
func wrapError(err error) error {
	if err == nil {
		return nil
	}
	// 同类型
	if se := new(errors.Error); errors.As(err, &se) {
		return err
	}
	return err
}
