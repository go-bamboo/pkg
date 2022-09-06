package jsonx

import (
	"encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
)

func WrapJsonError(err error) error {
	if err == nil {
		return nil
	}
	// 同类型
	if se := new(errors.Error); errors.As(err, &se) {
		return err
	}
	if se := new(json.SyntaxError); errors.As(err, &se) {
		return JsonSyntaxError(se)
	}
	return errors.FromError(err)
}

func JsonSyntaxError(err *json.SyntaxError) error {
	return errors.InternalServer("JsonSyntaxError", err.Error())
}

func IsJsonSyntaxError(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "JsonSyntaxError" && se.Code == 500
}
