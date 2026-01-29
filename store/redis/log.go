package redis

import (
	"context"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/log/sugar"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	// zap logger
	slogger *sugar.ZapLogger
	level   zapcore.Level
}

func NewLogger() *Logger {
	rlog := &Logger{
		level:   zapcore.DebugLevel, // 默认debug
		slogger: log.WithOpts(sugar.WithSkip(3)),
	}

	// copy
	return rlog
}

func (l Logger) Enabled(lvl zapcore.Level) bool {
	//if lvl >= l.level {
	//	return true
	//}
	//return false
	return true
}

func (l Logger) Printf(ctx context.Context, format string, v ...interface{}) {
	l.slogger.Debugf(format, v...)
}
