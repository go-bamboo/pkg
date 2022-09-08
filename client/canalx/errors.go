package canalx

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
)

// WrapPingcapError 转换错误
func WrapPingcapError(err error) error {
	if err == nil {
		return nil
	}
	// 同类型
	if se := new(errors.Error); errors.As(err, &se) {
		return err
	}
	return errors.InternalServer("PingErr", fmt.Sprintf("%+x", err))
}
