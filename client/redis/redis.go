package redis

import (
	"crypto/tls"
	"github.com/go-bamboo/pkg/log"
	"github.com/go-redis/redis/v8"
)

const KeepTTL = redis.KeepTTL

type (
	Pipeliner = redis.Pipeliner
	Script    = redis.Script
	Z         = redis.Z
	ZAddArgs  = redis.ZAddArgs
	ZRangeBy  = redis.ZRangeBy
	XAddArgs  = redis.XAddArgs
)

func NewScript(src string) *Script {
	return redis.NewScript(src)
}

type Client struct {
	redis.Client
}

type ClusterClient struct {
	redis.ClusterClient
}

func New(c *Conf) *Client {
	opts := &redis.Options{
		Addr:         c.Addrs[0],
		DialTimeout:  c.DialTimeout.AsDuration(),
		ReadTimeout:  c.ReadTimeout.AsDuration(),
		WriteTimeout: c.DialTimeout.AsDuration(),
		DB:           int(c.Db),
	}
	if len(c.Username) > 0 {
		opts.Username = c.Username
	}
	if len(c.Password) > 0 {
		opts.Password = c.Password
	}
	if c.Tls != nil && c.Tls.InsecureSkipVerify {
		opts.TLSConfig = &tls.Config{InsecureSkipVerify: c.Tls.InsecureSkipVerify}
	}
	rdb := redis.NewClient(opts)
	rdb.AddHook(NewRedisTracingHook(c.Debug))
	if c.Debug {
		redis.SetLogger(NewLogger(log.GetCore()))
	}
	return &Client{
		Client: *rdb,
	}
}

func NewCluster(c *Conf) *ClusterClient {
	opts := &redis.ClusterOptions{
		Addrs:        c.Addrs,
		DialTimeout:  c.DialTimeout.AsDuration(),
		ReadTimeout:  c.ReadTimeout.AsDuration(),
		WriteTimeout: c.DialTimeout.AsDuration(),
	}
	if len(c.Username) > 0 {
		opts.Username = c.Username
	}
	if len(c.Password) > 0 {
		opts.Password = c.Password
	}
	if c.Tls != nil && c.Tls.InsecureSkipVerify {
		opts.TLSConfig = &tls.Config{InsecureSkipVerify: c.Tls.InsecureSkipVerify}
	}
	rdb := redis.NewClusterClient(opts)
	rdb.AddHook(NewRedisTracingHook(c.Debug))
	if c.Debug {
		redis.SetLogger(NewLogger(log.GetCore()))
	}
	return &ClusterClient{
		ClusterClient: *rdb,
	}
}
