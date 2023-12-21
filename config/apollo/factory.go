package apollo

import (
	"github.com/go-bamboo/pkg/config"
	"github.com/go-bamboo/pkg/log"
	"gopkg.in/yaml.v3"
	"net/url"
)

func init() {
	config.Register("apollo", Create)
}

func Create(uri *url.URL, v interface{}) (config.Config, error) {
	q := uri.Query()
	appId := q.Get("appid")
	namespace := q.Get("ns")
	c := config.New(
		config.WithSource(
			NewConfigSource(
				AppID(appId),
				Namespaces(namespace+".yaml"),
				MetaAddr("http://"+uri.Host),
				SkipLocalCache(),
				WithLogger(log.GetLogger())),
		),
		config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
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
