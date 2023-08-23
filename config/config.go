package config

import (
	"github.com/go-kratos/kratos/v2/config"
	"net/url"
	"os"
)

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
	conf := os.Getenv("NACOS_ADDRESS")
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
