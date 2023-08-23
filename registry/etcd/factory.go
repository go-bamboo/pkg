package etcd

import (
	"github.com/go-bamboo/pkg/registry"
	etcdv3 "go.etcd.io/etcd/client/v3"
)

func init() {
	registry.Register("Etcd", Create)
}

func Create(c *registry.Conf) (registry.Registrar, registry.Discovery, error) {
	cli, err := etcdv3.New(etcdv3.Config{
		Endpoints: c.Endpoints,
	})
	if err != nil {
		panic(err)
	}
	r := New(cli)
	return r, r, nil
}
