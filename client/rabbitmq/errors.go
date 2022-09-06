package rabbitmq

import (
	"github.com/go-kratos/kratos/v2/errors"
)

func wrapError(err error) error {
	if err == nil {
		return nil
	}
	return errors.FromError(err)
}
