package config

import (
	"flag"
	"net/url"
	"path"

	"github.com/go-bamboo/pkg/apollo"
	"github.com/go-bamboo/pkg/filex"
	"github.com/go-bamboo/pkg/log"
	"github.com/go-kratos/kratos/contrib/config/consul/v2"
	nacos "github.com/go-kratos/kratos/contrib/config/nacos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/hashicorp/consul/api"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/yaml.v3"
)

type Value = config.Value

var (
	// conf is the config remote addr
	conf string
)

func init() {
	flag.StringVar(&conf, "conf", "file:///../../configs/conf.yaml", "url for config eg: file:///../../configs/conf.yaml")
}

func Load(v interface{}) config.Config {
	flag.Parse()
	uri, err := url.Parse(conf)
	if err != nil {
		panic(err)
	}
	if uri.Scheme == "file" {
		cp := filex.GetCurrentPath()
		c := config.New(
			config.WithSource(
				file.NewSource(path.Join(cp, uri.Path)),
			),
			config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
				return yaml.Unmarshal(kv.Value, v)
			}))
		if err := c.Load(); err != nil {
			panic(err)
		}
		if err := c.Scan(v); err != nil {
			panic(err)
		}
		return c
	} else if uri.Scheme == "apollo" {
		q := uri.Query()
		appId := q.Get("appid")
		namespace := q.Get("ns")
		c := config.New(
			config.WithSource(
				apollo.NewConfigSource(
					apollo.AppID(appId),
					apollo.Namespaces(namespace+".yaml"),
					apollo.MetaAddr("http://"+uri.Host),
					apollo.SkipLocalCache(),
					apollo.WithLogger(log.DefaultLogger)),
			),
			config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
				return yaml.Unmarshal(kv.Value, v)
			}))
		if err := c.Load(); err != nil {
			panic(err)
		}
		if err := c.Scan(v); err != nil {
			panic(err)
		}
		return c
	} else if uri.Scheme == "consul" {
		consulConfig := api.DefaultConfig()
		consulConfig.Address = uri.Host
		consulClient, err := api.NewClient(consulConfig)
		if err != nil {
			panic(err)
		}
		cs, err := consul.New(consulClient, consul.WithPath(uri.Path))
		if err != nil {
			panic(err)
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
		return c
	} else if uri.Scheme == "nacos" {
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
		source := nacos.NewConfigSource(client, nacos.WithGroup(group), nacos.WithDataID(dataId+".yaml"))
		c := config.New(
			config.WithSource(source),
			config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
				return yaml.Unmarshal(kv.Value, v)
			}))
		if err := c.Load(); err != nil {
			panic(err)
		}
		if err := c.Scan(v); err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}
	return nil
}
