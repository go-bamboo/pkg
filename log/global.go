package log

import (
	"context"
	"fmt"
	"runtime"
	"sync"

	"github.com/go-bamboo/pkg/log/std"
	"github.com/go-bamboo/pkg/log/sugar"
	"github.com/go-kratos/kratos/v2/errors"
	"go.uber.org/zap/zapcore"
)

// DefaultLogger is default logger.
var defaultLogger = std.NewStdCore(zapcore.DebugLevel)

func init() {
	SetLogger(sugar.NewLogger(defaultLogger, sugar.WithSkip(1)))
}

// globalLogger is designed as a global logger in current process.
var global = &loggerAppliance{}

// loggerAppliance is the proxy of `Logger` to
// make logger change will affect all sub-logger.
type loggerAppliance struct {
	lock sync.Mutex
	sugar.ZapLogger
}

func (a *loggerAppliance) SetLogger(in *sugar.ZapLogger) {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.ZapLogger = *in
}

func (a *loggerAppliance) GetLogger() *sugar.ZapLogger {
	return &a.ZapLogger
}

// SetLogger should be called before any other log call.
// And it is NOT THREAD SAFE.
func SetLogger(logger *sugar.ZapLogger) {
	global.SetLogger(logger)
}

// GetLogger returns global logger appliance as logger in current process.
func GetLogger() *sugar.ZapLogger {
	return global.GetLogger()
}

func With(kv ...interface{}) *sugar.ZapLogger {
	return global.GetLogger().With(kv...)
}

func WithOpts(opts ...sugar.Option) *sugar.ZapLogger {
	return global.GetLogger().WithOpts(opts...)
}

// Debug logs a message at debug level.
func Debug(a ...interface{}) {
	global.Debug(a...)
}

// Debugf logs a message at debug level.
func Debugf(format string, a ...interface{}) {
	global.Debugf(format, a...)
}

// Debugw logs a message at debug level.
func Debugw(msg string, keyvals ...interface{}) {
	global.Debugw(msg, keyvals...)
}

// DebugwCtx logs a message at debug level.
func DebugwCtx(ctx context.Context, msg string, keyvals ...interface{}) {
	global.DebugwCtx(ctx, msg, keyvals...)
}

// Info logs a message at info level.
func Info(a ...interface{}) {
	global.Info(a...)
}

// Infof logs a message at info level.
func Infof(format string, a ...interface{}) {
	global.Infof(format, a...)
}

// Infow logs a message at info level.
func Infow(msg string, keyvals ...interface{}) {
	global.Infow(msg, keyvals...)
}

// InfowCtx logs a message at info level.
func InfowCtx(ctx context.Context, msg string, keyvals ...interface{}) {
	global.InfowCtx(ctx, msg, keyvals...)
}

// Warn logs a message at warn level.
func Warn(a ...interface{}) {
	global.Warn(a...)
}

// Warnf logs a message at warnf level.
func Warnf(format string, a ...interface{}) {
	global.Warnf(format, a...)
}

// Warnw logs a message at warnf level.
func Warnw(msg string, keyvals ...interface{}) {
	global.Warnw(msg, keyvals...)
}

// WarnwCtx logs a message at warnf level.
func WarnwCtx(ctx context.Context, msg string, keyvals ...interface{}) {
	global.WarnwCtx(ctx, msg, keyvals...)
}

// Error logs a message at error level.
func Error(a ...interface{}) {
	global.Error(a...)
}

// Errorf logs a message at error level.
func Errorf(format string, a ...interface{}) {
	global.Errorf(format, a...)
}

// Errorw logs a message at error level.
func Errorw(msg string, keyvals ...interface{}) {
	global.Errorw(msg, keyvals...)
}

// ErrorwCtx logs a message at error level.
func ErrorwCtx(ctx context.Context, msg string, keyvals ...interface{}) {
	global.ErrorwCtx(ctx, msg, keyvals...)
}

// ErrorPanic logs a message at error level.
func ErrorPanic(err interface{}) {
	const size = 64 << 10
	buf := make([]byte, size)
	buf = buf[:runtime.Stack(buf, false)]
	pl := fmt.Sprintf("scan call panic: %v\n%s\n", err, buf)
	global.Errorf("%s", pl)
}

// ErrorStack logs a message at error level.
func ErrorStack(err error) {
	se := errors.FromError(err)
	if se == nil {
		global.Error(err)
		return
	}
	global.Errorw(fmt.Sprintf("%+v", err), "code", se.Code, "reason", se.Reason, "msg", se.Message, "md", se.Metadata)
}

// Fatal logs a message at fatal level.
func Fatal(a ...interface{}) {
	global.Fatal(a...)
}

// Fatalf logs a message at fatal level.
func Fatalf(format string, a ...interface{}) {
	global.Fatalf(format, a...)
}

// Fatalw logs a message at fatal level.
func Fatalw(msg string, keyvals ...interface{}) {
	global.Fatalw(msg, keyvals...)
}
