package cron

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	// zap logger
	logger  *zap.Logger
	slogger *zap.SugaredLogger
	level   zapcore.Level
}

func NewLogger(core zapcore.Core, lvl zapcore.Level) *Logger {

	rlog := &Logger{
		level: lvl, // 默认debug
	}

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	skip := zap.AddCallerSkip(1)

	// 构造日志
	logger := zap.New(core, caller, skip, zap.IncreaseLevel(rlog))
	slogger := logger.Sugar()
	rlog.logger = logger
	rlog.slogger = slogger

	// copy
	return rlog
}

func (l Logger) Enabled(lvl zapcore.Level) bool {
	if lvl >= l.level {
		return true
	}
	return false
}

func (s *Logger) Info(msg string, keysAndValues ...interface{}) {
	s.slogger.Infow(msg, keysAndValues...)
}

// Error logs an error condition.
func (s *Logger) Error(err error, msg string, keysAndValues ...interface{}) {
	if len(keysAndValues) == 0 {
		s.slogger.Error(msg)
		return
	}
	if len(keysAndValues)%2 != 0 {
		keysAndValues = append(keysAndValues, "")
	}
	keysAndValues = append(keysAndValues, "err", err)
	s.slogger.Errorw(msg, keysAndValues...)
}
