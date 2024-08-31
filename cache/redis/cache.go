package redis

import (
	"context"
	"time"

	"github.com/go-bamboo/pkg/cache"
	"github.com/go-bamboo/pkg/store/redis"
	redc "github.com/go-redis/cache/v8"
)

type redisCache struct {
	che *redc.Cache
}

func New(c *redis.Conf, localTTL time.Duration) cache.Cache {
	rdb := redis.New(c)
	mycache := redc.New(&redc.Options{
		Redis:      rdb,
		LocalCache: redc.NewTinyLFU(100000, localTTL),
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
	return val, time.Now(), nil
}

// Put stores a key-value pair into cache.
func (c redisCache) Put(ctx context.Context, key string, val interface{}, ttl time.Duration) error {
	return c.che.Set(&redc.Item{
		Ctx:   ctx,
		Key:   key,
		Value: val,
		TTL:   ttl,
	})
}

// Delete removes a key from cache.
func (c redisCache) Delete(ctx context.Context, key string) error {
	return c.che.Delete(ctx, key)
}

// String returns the name of the implementation.
func (c redisCache) String() string {
	return "redis"
}
