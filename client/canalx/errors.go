package canalx

import (
	"bls/pkg/ecode"
	"github.com/go-kratos/kratos/v2/errors"
	pingcaperr "github.com/pingcap/errors"
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
	if se := new(pingcaperr.Error); errors.As(err, &se) {
		return errors.InternalServer(string(se.ID()), se.GetMsg())
	}
	return ecode.Unknown(err.Error())
}
