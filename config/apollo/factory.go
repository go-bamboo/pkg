package apollo

import (
	"github.com/go-bamboo/pkg/config"
	"github.com/go-bamboo/pkg/log"
	configx "github.com/go-kratos/kratos/v2/config"
	"gopkg.in/yaml.v3"
	"net/url"
)

func init() {
	config.Register("apollo", Create)
}

func Create(uri *url.URL, v interface{}) (configx.Config, error) {
	q := uri.Query()
	appId := q.Get("appid")
	namespace := q.Get("ns")
	c := configx.New(
		configx.WithSource(
			NewConfigSource(
				AppID(appId),
				Namespaces(namespace+".yaml"),
				MetaAddr("http://"+uri.Host),
				SkipLocalCache(),
				WithLogger(log.DefaultLogger)),
		),
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
