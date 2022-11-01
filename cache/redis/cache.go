package redis

import (
	"context"
	"github.com/go-bamboo/pkg/cache"
	"github.com/go-bamboo/pkg/client/redis"
	redc "github.com/go-redis/cache/v8"
	"time"
)

type redisCache struct {
	che *redc.Cache
}

func New(c *redis.Conf) cache.Cache {
	rdb := redis.New(c)
	mycache := redc.New(&redc.Options{
		Redis:      rdb,
		LocalCache: redc.NewTinyLFU(1000, time.Minute),
	})
	return &redisCache{
		che: mycache,
	}
}

func (c redisCache) Get(ctx context.Context, key string) (interface{}, time.Time, error) {
	var val interface{}
	if err := c.che.Get(ctx, key, val); err != nil {
		return nil, time.Now(), err
	}
	return nil, time.Now(), nil
}

// Put stores a key-value pair into cache.
func (c redisCache) Put(ctx context.Context, key string, val interface{}, d time.Duration) error {
	return nil
}

// Delete removes a key from cache.
func (c redisCache) Delete(ctx context.Context, key string) error {
	return nil
}

// String returns the name of the implementation.
func (c redisCache) String() string {
	return "redis"
}
