package fluent

import (
	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/log/core"
	"go.uber.org/zap/zapcore"
)

func init() {
	log.Register("fluent", Create)
}

func Create(c *log.Conf) (core.Logger, error) {
	cc, err := NewFluentCore(
		Level(zapcore.Level(c.Level)),
		WithAddr(c.Endpoint),
	)
	if err != nil {
		return nil, err
	}
	return cc, nil
}
