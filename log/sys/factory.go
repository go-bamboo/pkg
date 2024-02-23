package sys

import (
	"github.com/go-bamboo/pkg/log/core"
)

func init() {
	core.Register("Sys", Create)
}

func Create(c *core.Conf) (core.Logger, error) {
	return nil, nil
}
