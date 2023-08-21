package config

import (
	"flag"
	"github.com/go-kratos/kratos/v2/config"
	"net/url"
	"os"
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
	nacosAddr := os.Getenv("NACOS_ADDRESS")
	if len(nacosAddr) > 0 {
		conf = nacosAddr
	}
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
