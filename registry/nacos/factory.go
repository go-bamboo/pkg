package nacos

import (
	"github.com/go-bamboo/pkg/registry"
	"github.com/go-bamboo/pkg/registry/core"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func init() {
	registry.Register("kube", Create)
}

func Create(c *registry.Conf) (core.Registrar, core.Discovery, error) {
	sc := []constant.ServerConfig{}
	for _, server := range c.Nacos.Servers {
		sc = append(sc, *constant.NewServerConfig(server.IpAddr, server.Port))
	}
	cc := constant.ClientConfig{
		NamespaceId:         c.Nacos.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              c.Nacos.LogDir,
		CacheDir:            c.Nacos.CacheDir,
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
