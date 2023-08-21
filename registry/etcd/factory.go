package etcd

import (
	"github.com/go-bamboo/pkg/registry"
	"github.com/go-bamboo/pkg/registry/core"
	etcdv3 "go.etcd.io/etcd/client/v3"
)

func init() {
	registry.Register("etcd", Create)
}

func Create(c *registry.Conf) (core.Registrar, core.Discovery, error) {
	cli, err := etcdv3.New(etcdv3.Config{
		Endpoints: c.Etcd.Endpoints,
	})
	if err != nil {
		panic(err)
	}
	r := New(cli)
	return r, r, nil
}
