package apollo

import (
	"net/url"

	"github.com/go-bamboo/pkg/config"
	"github.com/go-bamboo/pkg/log"
	"github.com/go-kratos/kratos/v2/encoding"
)

func init() {
	config.Register("apollo", Create)
}

func Create(uri *url.URL, v interface{}, format string) (config.Config, error) {
	q := uri.Query()
	appId := q.Get("appid")
	namespace := q.Get("ns")
	c := config.New(
		config.WithSource(
			NewConfigSource(
				AppID(appId),
				Namespaces(namespace+"."+format),
				MetaAddr("http://"+uri.Host),
				SkipLocalCache(),
				WithLogger(log.GetLogger()),
				WithFormat(format),
			),
		),
		config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
			return encoding.GetCodec(kv.Format).Unmarshal(kv.Value, v)
		}))
	if err := c.Load(); err != nil {
		return nil, err
	}
	if err := c.Scan(v); err != nil {
		return nil, err
	}
	return c, nil
}
