package log

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/go-bamboo/pkg/log/std"
	"github.com/go-kratos/kratos/v2/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// DefaultLogger is default logger.
var DefaultLogger = NewLogger(std.NewStdCore(zapcore.DebugLevel), 1)

// globalLogger is designed as a global logger in current process.
var global = &loggerAppliance{}

// loggerAppliance is the proxy of `Logger` to
// make logger change will affect all sub-logger.
type loggerAppliance struct {
	lock sync.Mutex
	*ZapLogger
}

func init() {
	global.SetLogger(DefaultLogger)
}

func (a *loggerAppliance) SetLogger(in *ZapLogger) {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.ZapLogger = in
}

func (a *loggerAppliance) GetLogger() *zap.SugaredLogger {
	return a.slogger
}

// SetLogger should be called before any other log call.
// And it is NOT THREAD SAFE.
func SetLogger(logger *ZapLogger) {
	global.SetLogger(logger)
}

// GetLogger returns global logger appliance as logger in current process.
func GetLogger() *ZapLogger {
	return global.ZapLogger
}

func GetCore() zapcore.Core {
	return global.ZapLogger.logger.Core()
}

// Log Print log by level and keyvals.
//func Log(level Level, keyvals ...interface{}) {
//	global.slogger.l(level, keyvals...)
//}

func With(kv ...interface{}) *ZapLogger {
	core := global.ZapLogger.logger.Core()
	core = WithCore(core, kv...)
	return NewLogger(core, 1)
}

// Debug logs a message at debug level.
func Debug(a ...interface{}) {
	global.slogger.Debug(a...)
}

// Debugf logs a message at debug level.
func Debugf(format string, a ...interface{}) {
	global.slogger.Debugf(format, a...)
}

// Debugw logs a message at debug level.
func Debugw(msg string, keyvals ...interface{}) {
	global.slogger.Debugw(msg, keyvals...)
}

// Info logs a message at info level.
func Info(a ...interface{}) {
	global.slogger.Info(a...)
}

// Infof logs a message at info level.
func Infof(format string, a ...interface{}) {
	global.slogger.Infof(format, a...)
}

// Infow logs a message at info level.
func Infow(msg string, keyvals ...interface{}) {
	global.slogger.Infow(msg, keyvals...)
}

// Warn logs a message at warn level.
func Warn(a ...interface{}) {
	global.slogger.Warn(a...)
}

// Warnf logs a message at warnf level.
func Warnf(format string, a ...interface{}) {
	global.slogger.Warnf(format, a...)
}

// Warnw logs a message at warnf level.
func Warnw(msg string, keyvals ...interface{}) {
	global.slogger.Warnw(msg, keyvals...)
}

// Error logs a message at error level.
func Error(a ...interface{}) {
	global.slogger.Error(a...)
}

// Errorf logs a message at error level.
func Errorf(format string, a ...interface{}) {
	global.slogger.Errorf(format, a...)
}

// Errorw logs a message at error level.
func Errorw(msg string, keyvals ...interface{}) {
	global.slogger.Errorw(msg, keyvals...)
}

// ErrorPanic logs a message at error level.
func ErrorPanic(err interface{}) {
	const size = 64 << 10
	buf := make([]byte, size)
	buf = buf[:runtime.Stack(buf, false)]
	pl := fmt.Sprintf("scan call panic: %v\n%s\n", err, buf)
	global.slogger.Errorf("%s", pl)
}

// ErrorStack logs a message at error level.
func ErrorStack(err error) {
	se := errors.FromError(err)
	if se == nil {
		global.slogger.Error(err)
		return
	}
	global.slogger.Errorw(fmt.Sprintf("%+v", err), "code", se.Code, "reason", se.Reason, "msg", se.Message, "md", se.Metadata)
}

// Fatal logs a message at fatal level.
func Fatal(a ...interface{}) {
	global.slogger.Fatal(a...)
}

// Fatalf logs a message at fatal level.
func Fatalf(format string, a ...interface{}) {
	global.slogger.Fatalf(format, a...)
}

// Fatalw logs a message at fatal level.
func Fatalw(msg string, keyvals ...interface{}) {
	global.slogger.Fatalw(msg, keyvals...)
}
