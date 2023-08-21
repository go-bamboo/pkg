package nacos

import (
	"github.com/go-bamboo/pkg/config"
	configx "github.com/go-kratos/kratos/v2/config"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/yaml.v3"
	"net/url"
)

func init() {
	config.Register("file", Create)
}

func Create(uri *url.URL, v interface{}) (configx.Config, error) {
	q := uri.Query()
	namespace := q.Get("namespace")
	group := q.Get("group")
	dataId := q.Get("dataId")
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(uri.Host, 80),
	}
	cc := &constant.ClientConfig{
		NamespaceId:         namespace, //namespace id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	// a more graceful way to create naming client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}
	source := NewConfigSource(client, WithGroup(group), WithDataID(dataId+".yaml"))
	c := configx.New(
		configx.WithSource(source),
		configx.WithDecoder(func(kv *configx.KeyValue, v map[string]interface{}) error {
			return yaml.Unmarshal(kv.Value, v)
		}))
	if err := c.Load(); err != nil {
		return nil, err
	}
	if err := c.Scan(v); err != nil {
		return nil, err
	}
	return c, nil
}
