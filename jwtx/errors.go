package jwtx

import (
	"fmt"

	"github.com/emberfarkas/pkg/ecode"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/golang-jwt/jwt/v4"
)

func WrapError(err error) error {
	if err == nil {
		return nil
	}
	if se := new(errors.Error); errors.As(err, &se) {
		return err
	} else if ve := new(jwt.ValidationError); errors.As(err, &ve) {
		if ve.Errors == jwt.ValidationErrorAudience {
			return ValidationErrorExpired(err)
		} else if ve.Errors == jwt.ValidationErrorExpired {
			return ValidationErrorExpired(err)
		} else {
			return errors.InternalServer(fmt.Sprint(ve.Errors), ve.Error())
		}
	} else if errors.Is(err, jwt.ErrInvalidKey) {
		return ErrInvalidKey(err)
	} else {
		return ecode.Unknown(err.Error())
	}
}

func ValidationErrorAudience(err error) error {
	return errors.InternalServer("ValidationErrorAudience", err.Error())
}

func IsValidationErrorAudience(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "ValidationErrorAudience" && se.Code == 500
}

func ValidationErrorExpired(err error) error {
	return errors.InternalServer("ValidationErrorExpired", err.Error())
}

func IsValidationErrorExpired(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "ValidationErrorExpired" && se.Code == 500
}

func ErrInvalidKey(err error) error {
	return errors.InternalServer("ErrInvalidKey", err.Error())
}

func IsErrInvalidKey(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "ErrInvalidKey" && se.Code == 500
}
