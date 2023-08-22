package nacos

import (
	"github.com/go-bamboo/pkg/registry"
	"github.com/go-bamboo/pkg/registry/core"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"net/url"
	"strconv"
)

func init() {
	registry.Register("Nacos", Create)
}

func Create(c *registry.Conf) (core.Registrar, core.Discovery, error) {
	var sc []constant.ServerConfig
	for _, server := range c.Endpoints {
		uri, err := url.Parse(server)
		if err != nil {
			return nil, nil, err
		}
		port, err := strconv.ParseUint(uri.Port(), 10, 64)
		if err != nil {
			return nil, nil, err
		}
		sc = append(sc, *constant.NewServerConfig(uri.Host, port))
	}
	cc := constant.ClientConfig{
		NamespaceId:         c.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              c.LogDir,
		CacheDir:            c.CacheDir,
		LogLevel:            "info",
	}
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ServerConfigs: sc,
			ClientConfig:  &cc,
		},
	)
	if err != nil {
		panic(err)
	}
	r := New(client)
	return r, r, nil
}
