package file

import (
	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/log/core"
	"go.uber.org/zap/zapcore"
)

func init() {
	log.Register("file", Create)
}

func Create(c *log.Conf) (core.Logger, error) {
	return NewFileCore(
		Level(zapcore.Level(c.Level)),
		WithPath(c.Path),
		WithName(c.LogGroupName),
	), nil
}
