package sra

type RedisLimiter struct {
}

func NewRedisLimiter() *RedisLimiter {
	return NewRedisLimiter()
}

func (l *RedisLimiter) Allow() error {
	return nil
}

func (l *RedisLimiter) ReportResult(result error) {
	return
}
