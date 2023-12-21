package consul

import (
	"github.com/go-bamboo/pkg/config"
	"github.com/hashicorp/consul/api"
	"gopkg.in/yaml.v3"
	"net/url"
)

func init() {
	config.Register("file", Create)
}

func Create(uri *url.URL, v interface{}) (config.Config, error) {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = uri.Host
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		panic(err)
	}
	cs, err := New(consulClient, WithPath(uri.Path))
	if err != nil {
		return nil, err
	}
	c := config.New(
		config.WithSource(cs),
		config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
			return yaml.Unmarshal(kv.Value, v)
		}))
	if err := c.Load(); err != nil {
		panic(err)
	}
	if err := c.Scan(v); err != nil {
		panic(err)
	}
	return c, nil
}
