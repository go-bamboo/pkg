package core

import (
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	zapcore.Core
	Close()
}
