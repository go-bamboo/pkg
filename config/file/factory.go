package file

import (
	"fmt"
	"net/url"
	"path"

	"github.com/go-bamboo/pkg/config"
	"github.com/go-bamboo/pkg/filex"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/encoding"
)

func init() {
	config.Register("file", Create)
}

func Create(uri *url.URL, v interface{}, format string) (config.Config, error) {
	cp := filex.GetCurrentPath()
	c := config.New(
		config.WithSource(
			file.NewSource(path.Join(cp, uri.Path)),
		),
		config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
			if format == "" {
				format = kv.Format
			}
			fmt.Println(kv.Key, kv.Format, format)
			if err := encoding.GetCodec(format).Unmarshal(kv.Value, v); err != nil {
				return err
			}
			return nil
		}))
	if err := c.Load(); err != nil {
		panic(err)
	}
	if err := c.Scan(v); err != nil {
		panic(err)
	}
	return c, nil
}
