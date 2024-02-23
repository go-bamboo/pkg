package file

import (
	"github.com/go-bamboo/pkg/log/core"
	"go.uber.org/zap/zapcore"
)

func init() {
	core.Register("File", Create)
}

func Create(c *core.Conf) (core.Logger, error) {
	return NewFileCore(
		Level(zapcore.Level(c.Level)),
		WithPath(c.Path),
		WithName(c.LogGroupName),
	), nil
}
