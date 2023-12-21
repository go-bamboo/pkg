package file

import (
	"github.com/go-bamboo/pkg/config"
	"github.com/go-bamboo/pkg/filex"
	"github.com/go-kratos/kratos/v2/config/file"
	"gopkg.in/yaml.v3"
	"net/url"
	"path"
)

func init() {
	config.Register("file", Create)
}

func Create(uri *url.URL, v interface{}) (config.Config, error) {
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
	return c, nil
}
