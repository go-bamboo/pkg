package redis

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-redis/redis/v8"
)

func wrapRedisError(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, redis.Nil) {
		return RedisNil(redis.Nil)
	}
	return err
}

func ErrSpanLost(format string, a ...interface{}) error {
	return errors.InternalServer("ErrSpanLost", fmt.Sprintf(format, a...))
}

func IsErrSpanLost(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "ErrSpanLost" && se.Code == 500
}

// RedisNil 没有值
func RedisNil(err redis.Error) error {
	return errors.InternalServer("RedisNil", err.Error())
}

func IsRedisNil(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "RedisNil" && se.Code == 500
}

func IsNil(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "RedisNil" && se.Code == 500
}
