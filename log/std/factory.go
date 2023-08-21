package std

import (
	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/log/core"
	"go.uber.org/zap/zapcore"
)

func init() {
	log.Register("Stdout", Create)
}

func Create(c *log.Conf) (core.Logger, error) {
	return NewStdCore(zapcore.Level(c.Level)), nil
}
