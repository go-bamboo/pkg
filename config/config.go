package config

import (
	"flag"
	"github.com/go-kratos/kratos/v2/config"
	"net/url"
	"os"
)

var (
	// conf is the config remote addr
	conf string
)

const KEY_NACOS = "NACOS_ADDRESS"

func init() {
	flag.StringVar(&conf, "conf", "file:///../../configs/conf.yaml", "url for config eg: file:///../../configs/conf.yaml")
}

type Value = config.Value

func Load(conf string, v interface{}) config.Config {
	uri, err := url.Parse(conf)
	if err != nil {
		panic(err)
	}
	c, err := Create(uri, v)
	if err != nil {
		panic(err)
	}
	return c
}

func LoadEnv(v interface{}) config.Config {
	conf := os.Getenv(KEY_NACOS)
	uri, err := url.Parse(conf)
	if err != nil {
		panic(err)
	}
	c, err := Create(uri, v)
	if err != nil {
		panic(err)
	}
	return c
}

func LoadFlag(v interface{}) config.Config {
	flag.Parse()
	uri, err := url.Parse(conf)
	if err != nil {
		panic(err)
	}
	c, err := Create(uri, v)
	if err != nil {
		panic(err)
	}
	return c
}
