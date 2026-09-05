package redis

import (
	"fmt"

	"github.com/go-bamboo/pkg/registry"
	"github.com/go-bamboo/pkg/store/redis"
)

func init() {
	registry.Register("Redis", Create)
}

func Create(c *registry.Conf) (registry.Registrar, registry.Discovery, error) {
	if c == nil || len(c.Endpoints) == 0 {
		return nil, nil, fmt.Errorf("redis registry: endpoints is empty")
	}
	rdb := redis.New(&redis.Conf{
		Addrs:       c.Endpoints,
		DialTimeout: c.DialTimeout,
	})
	ropts := []Option{WithHealthCheck(c.EnableHealthCheck)}
	if c.Namespace != "" {
		ropts = append(ropts, WithNamespace(c.Namespace))
	}
	r := New(rdb, ropts...)
	return r, r, nil
}
