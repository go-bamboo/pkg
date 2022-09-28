package redis

import (
	"time"

	"github.com/go-redis/cache/v8"
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

	che cache.Cache
}

func New(c *Conf) *Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Addr,
		DialTimeout:  c.DialTimeout.AsDuration(),
		ReadTimeout:  c.ReadTimeout.AsDuration(),
		WriteTimeout: c.DialTimeout.AsDuration(),
		//Username:     c.Redis.Username,
		Password: c.Password,
		DB:       int(c.Db),
	})
	rdb.AddHook(NewRedisTracingHook())

	mycache := cache.New(&cache.Options{
		Redis:      rdb,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	return &Client{
		Client: *rdb,
		che:    *mycache,
	}
}
