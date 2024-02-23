package std

import (
	"github.com/go-bamboo/pkg/log/core"
	"go.uber.org/zap/zapcore"
)

func init() {
	core.Register("Stdout", Create)
}

func Create(c *core.Conf) (core.Logger, error) {
	return NewStdCore(zapcore.Level(c.Level)), nil
}
