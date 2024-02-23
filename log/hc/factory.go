package hc

import (
	"github.com/go-bamboo/pkg/log/core"
)

func init() {
	core.Register("hc", Create)
}

func Create(c *core.Conf) (core.Logger, error) {
	return nil, nil
}
