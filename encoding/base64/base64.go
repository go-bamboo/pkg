package base64

import (
	"encoding/base64"
	"github.com/go-kratos/kratos/v2/errors"
)

func DecodeString(s string) ([]byte, error) {
	storeKey, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		err = errors.FromError(err)
		return nil, err
	}
	return storeKey, err
}
