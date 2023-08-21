package std

import (
	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/log/core"
	"go.uber.org/zap/zapcore"
)

// DefaultLogger is default logger.
var DefaultLogger = NewStdCore(zapcore.DebugLevel)

func init() {
	log.Register("Stdout", Create)
	log.SetLogger(log.NewLogger(DefaultLogger, 1))
}

func Create(c *log.Conf) (core.Logger, error) {
	return NewStdCore(zapcore.Level(c.Level)), nil
}
