package migrate

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	// zap logger
	logger  *zap.Logger
	slogger *zap.SugaredLogger
	level   zapcore.Level
	verbose bool
}

func NewLogger(lvl zapcore.Level, core zapcore.Core) *Logger {

	rlog := &Logger{
		level:   lvl,
		verbose: false,
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

// Printf is like fmt.Printf
func (l Logger) Printf(format string, v ...interface{}) {
	l.slogger.Debugf(format, v...)
}

// Verbose should return true when verbose logging output is wanted
func (l Logger) Verbose() bool {
	return l.verbose
}
