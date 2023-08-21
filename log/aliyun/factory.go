package aliyun

import (
	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/log/core"
)

func init() {
	log.Register("AliYun", Create)
}

func Create(c *log.Conf) (core.Logger, error) {
	return nil, nil
}
