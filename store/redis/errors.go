package redis

//func wrapRedisError(err error) error {
//	if err == nil {
//		return nil
//	}
//	if errors.Is(err, redis.Nil) {
//		return RedisNil(redis.Nil)
//	}
//	return err
//}
//
//func ErrSpanLost(format string, a ...interface{}) error {
//	se := errors.InternalServer("ErrSpanLost", fmt.Sprintf(format, a...))
//	return se
//}
//
//func IsErrSpanLost(err error) bool {
//	se := errors.FromError(err)
//	return se.Reason == "ErrSpanLost" && se.Code == 500
//}
//
//// RedisNil 没有值
//func RedisNil(err redis.Error) error {
//	se := errors.InternalServer("RedisNil", err.Error())
//	se.WithCause(err)
//	return se
//}
//
//func IsRedisNil(err error) bool {
//	se := errors.FromError(err)
//	return se.Reason == "RedisNil" && se.Code == 500
//}
//
//func IsNil(err error) bool {
//	se := errors.FromError(err)
//	return se.Reason == "RedisNil" && se.Code == 500
//}
