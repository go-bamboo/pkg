package eth

import (
	"fmt"
	"strings"

	"github.com/go-bamboo/pkg/ecode"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/onrik/ethrpc"
)

// WrapError 转换错误
func WrapError(err error) error {
	if err == nil {
		return nil
	}
	// 同类型
	if se := new(errors.Error); errors.As(err, &se) {
		return err
	}
	if se := (ethrpc.EthError{}); errors.As(err, &se) {
		return errors.InternalServer(fmt.Sprintf("EthError:%v", se.Code), se.Message)
	}
	return ecode.Unknown(err.Error())
}

func IsNonceTooLow(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "EthError:-32000" && strings.Contains(se.Message, "nonce too low") && se.Code == 500
}
