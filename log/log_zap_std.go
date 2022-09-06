package log

import (
	"go.uber.org/zap/zapcore"
)

type ZapLoggerEx struct {
	ZapLogger
}

func NewZapLoggerEx(core zapcore.Core) *ZapLoggerEx {
	logger := NewLogger(core, 1)
	l := &ZapLoggerEx{
		ZapLogger: *logger,
	}
	return l
}

func (s *ZapLoggerEx) Info(msg string, keysAndValues ...interface{}) {
	s.slogger.Infow(msg, keysAndValues...)
}

// Error logs an error condition.
func (s *ZapLoggerEx) Error(err error, msg string, keysAndValues ...interface{}) {
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
