package rocketmq

import (
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type RocketLogger struct {
	// zap logger
	logger  *zap.Logger
	slogger *zap.SugaredLogger
	level   string
}

func NewLogger(core zapcore.Core) rlog.Logger {
	rlog := &RocketLogger{}

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

func (l RocketLogger) Enabled(lvl zapcore.Level) bool {
	if lvl >= zapcore.InfoLevel {
		return true
	}
	return false
}

func (l RocketLogger) Debug(msg string, fields map[string]interface{}) {
	var keysAndValues []zap.Field = l.cov(fields)
	l.logger.Debug(msg, keysAndValues...)
}

func (l RocketLogger) Info(msg string, fields map[string]interface{}) {
	var keysAndValues []zap.Field = l.cov(fields)
	l.logger.Debug(msg, keysAndValues...)
}
func (l RocketLogger) Warning(msg string, fields map[string]interface{}) {
	var keysAndValues []zap.Field = l.cov(fields)
	l.logger.Debug(msg, keysAndValues...)
}
func (l RocketLogger) Error(msg string, fields map[string]interface{}) {
	var keysAndValues []zap.Field = l.cov(fields)
	l.logger.Debug(msg, keysAndValues...)
}
func (l RocketLogger) Fatal(msg string, fields map[string]interface{}) {
	var keysAndValues []zap.Field = l.cov(fields)
	l.logger.Debug(msg, keysAndValues...)
}

func (l *RocketLogger) Level(level string) {
	l.level = level
}

func (l RocketLogger) OutputPath(path string) (err error) {
	return nil
}

func (l RocketLogger) cov(fields map[string]interface{}) []zap.Field {
	var keysAndValues []zap.Field
	for k, v := range fields {
		switch val := v.(type) {
		case bool:
			keysAndValues = append(keysAndValues, zap.Bool(k, val))
		case int:
			keysAndValues = append(keysAndValues, zap.Int(k, val))
		case int8:
			keysAndValues = append(keysAndValues, zap.Int8(k, val))
		case int16:
			keysAndValues = append(keysAndValues, zap.Int16(k, val))
		case int32:
			keysAndValues = append(keysAndValues, zap.Int32(k, val))
		case int64:
			keysAndValues = append(keysAndValues, zap.Int64(k, val))
		case uint:
			keysAndValues = append(keysAndValues, zap.Uint(k, val))
		case uint8:
			keysAndValues = append(keysAndValues, zap.Uint8(k, val))
		case uint16:
			keysAndValues = append(keysAndValues, zap.Uint16(k, val))
		case uint32:
			keysAndValues = append(keysAndValues, zap.Uint32(k, val))
		case uint64:
			keysAndValues = append(keysAndValues, zap.Uint64(k, val))
		case string:
			keysAndValues = append(keysAndValues, zap.String(k, val))
		}
	}
	return keysAndValues
}
