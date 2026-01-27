package config

import (
	"flag"
	"net/url"
	"os"

	"github.com/go-kratos/kratos/v2/config"
)

var (
	// conf is the config remote addr
	conf string
)

const KEY_NACOS = "NACOS_ADDRESS"

func init() {
	defaultConf := os.Getenv("APP_CONF")
	if defaultConf == "" {
		defaultConf = "file:///../../configs/conf.yaml" // 默认值
	}
	flag.StringVar(&conf, "conf", defaultConf, "url for config eg: file:///../../configs/conf.yaml")
}

type Value = config.Value
type Source = config.Source
type KeyValue = config.KeyValue
type Watcher = config.Watcher
type Config = config.Config

var New = config.New
var WithSource = config.WithSource
var WithDecoder = config.WithDecoder

func Load(conf string, v interface{}, format string) config.Config {
	uri, err := url.Parse(conf)
	if err != nil {
		panic(err)
	}
	c, err := Create(uri, v, format)
	if err != nil {
		panic(err)
	}
	return c
}

func LoadEnv(v interface{}, format string) config.Config {
	conf := os.Getenv(KEY_NACOS)
	uri, err := url.Parse(conf)
	if err != nil {
		panic(err)
	}
	c, err := Create(uri, v, format)
	if err != nil {
		panic(err)
	}
	return c
}

func LoadFlag(v interface{}, format string) config.Config {
	flag.Parse()
	uri, err := url.Parse(conf)
	if err != nil {
		panic(err)
	}
	c, err := Create(uri, v, format)
	if err != nil {
		panic(err)
	}
	return c
}
