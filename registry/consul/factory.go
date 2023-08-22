package consul

import (
	"github.com/go-bamboo/pkg/registry"
	"github.com/go-bamboo/pkg/registry/core"
	"github.com/hashicorp/consul/api"
)

func init() {
	registry.Register("Consul", Create)
}

func Create(c *registry.Conf) (core.Registrar, core.Discovery, error) {
	// consul
	consulConfig := api.DefaultConfig()
	consulConfig.Address = c.Endpoints[0]
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		return nil, nil, err
	}
	r := New(consulClient)
	return r, r, nil
}
