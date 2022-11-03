package redis

import (
	"crypto/tls"
	"github.com/go-redis/redis/v8"
)

const KeepTTL = redis.KeepTTL

type (
	Pipeliner = redis.Pipeliner
	Script    = redis.Script
	Z         = redis.Z
	ZAddArgs  = redis.ZAddArgs
	ZRangeBy  = redis.ZRangeBy
)

func NewScript(src string) *Script {
	return redis.NewScript(src)
}

type Client struct {
	redis.Client
}

func New(c *Conf) *Client {
	opts := &redis.Options{
		Addr:         c.Addr,
		DialTimeout:  c.DialTimeout.AsDuration(),
		ReadTimeout:  c.ReadTimeout.AsDuration(),
		WriteTimeout: c.DialTimeout.AsDuration(),
		//Username:     c.Redis.Username,
		Password: c.Password,
		DB:       int(c.Db),
	}
	if c.Tls != nil && c.Tls.InsecureSkipVerify {
		opts.TLSConfig = &tls.Config{InsecureSkipVerify: c.Tls.InsecureSkipVerify}
	}
	rdb := redis.NewClient(opts)
	rdb.AddHook(NewRedisTracingHook())
	return &Client{
		Client: *rdb,
	}
}
