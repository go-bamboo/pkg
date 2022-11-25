package redis

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	// zap logger
	logger  *zap.Logger
	slogger *zap.SugaredLogger
	level   zapcore.Level
}

func NewLogger(core zapcore.Core) *Logger {

	rlog := &Logger{
		level: zapcore.DebugLevel, // 默认debug
	}

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	skip := zap.AddCallerSkip(2)

	// 构造日志
	logger := zap.New(core, caller, skip, zap.IncreaseLevel(rlog))
	slogger := logger.Sugar()
	rlog.logger = logger
	rlog.slogger = slogger

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
