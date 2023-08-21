package sys

import (
	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/log/core"
)

func init() {
	log.Register("Sys", Create)
}

func Create(c *log.Conf) (core.Logger, error) {
	return nil, nil
}
