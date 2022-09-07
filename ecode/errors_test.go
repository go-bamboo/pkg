package ecode

import (
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/onrik/ethrpc"
)

func TestWrapError(t *testing.T) {
	se := new(ethrpc.EthError)
	se.Code = 301
	se.Message = "xx"
	err := WrapError(se)
	xe := errors.FromError(err)
	t.Errorf("code : %v", xe.Code)
	t.Errorf("message: %v", xe.Message)
}
