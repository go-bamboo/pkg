package tee

import (
	"time"

	"github.com/emberfarkas/pkg/log/core"
	"github.com/emberfarkas/pkg/log/file"
	"github.com/emberfarkas/pkg/log/fluent"
	"go.uber.org/zap/zapcore"
)

type Option func(*options)

type options struct {
	l      zapcore.Level
	stdout bool
	name   string

	// file
	path string

	// fluent
	addr               string
	timeout            time.Duration
	writeTimeout       time.Duration
	bufferLimit        int
	retryWait          int
	maxRetry           int
	maxRetryWait       int
	tagPrefix          string
	async              bool
	forceStopAsyncSend bool

	// hc
}

func Level(l zapcore.Level) Option {
	return func(c *options) {
		c.l = l
	}
}

func WithName(name string) Option {
	return func(c *options) {
		c.name = name
	}
}

func Stdout(stdout bool) Option {
	return func(c *options) {
		c.stdout = stdout
	}
}

func Path(path string) Option {
	return func(c *options) {
		c.path = path
	}
}

// WithAddr with config fluent addr.
func WithAddr(addr string) Option {
	return func(opts *options) {
		opts.addr = addr
	}
}

// WithTimeout with config Timeout.
func WithTimeout(timeout time.Duration) Option {
	return func(opts *options) {
		opts.timeout = timeout
	}
}

// WithWriteTimeout with config WriteTimeout.
func WithWriteTimeout(writeTimeout time.Duration) Option {
	return func(opts *options) {
		opts.writeTimeout = writeTimeout
	}
}

// WithBufferLimit with config BufferLimit.
func WithBufferLimit(bufferLimit int) Option {
	return func(opts *options) {
		opts.bufferLimit = bufferLimit
	}
}

// WithRetryWait with config RetryWait.
func WithRetryWait(retryWait int) Option {
	return func(opts *options) {
		opts.retryWait = retryWait
	}
}

// WithMaxRetry with config MaxRetry.
func WithMaxRetry(maxRetry int) Option {
	return func(opts *options) {
		opts.maxRetry = maxRetry
	}
}

// WithMaxRetryWait with config MaxRetryWait.
func WithMaxRetryWait(maxRetryWait int) Option {
	return func(opts *options) {
		opts.maxRetryWait = maxRetryWait
	}
}

// WithTagPrefix with config TagPrefix.
func WithTagPrefix(tagPrefix string) Option {
	return func(opts *options) {
		opts.tagPrefix = tagPrefix
	}
}

// WithAsync with config Async.
func WithAsync(async bool) Option {
	return func(opts *options) {
		opts.async = async
	}
}

// WithForceStopAsyncSend with config ForceStopAsyncSend.
func WithForceStopAsyncSend(forceStopAsyncSend bool) Option {
	return func(opts *options) {
		opts.forceStopAsyncSend = forceStopAsyncSend
	}
}

type teeCore struct {
	opts options

	hooks []core.Logger
	core  zapcore.Core
}

// NewTeeCore new a zap logger with options.
func NewTeeCore(opts ...Option) (core.Logger, error) {
	_options := options{
		l:      zapcore.DebugLevel,
		stdout: false,
	}
	for _, o := range opts {
		o(&_options)
	}

	// 设置日志级别
	level := _options.l
	cores := make([]zapcore.Core, 0)
	hooks := make([]core.Logger, 0)

	if _options.stdout {
		c := NewStdCore(level)
		cores = append(cores, c)
		hooks = append(hooks, c)
	}
	// file
	if len(_options.path) > 0 {
		// 每个日志级对应一个文件
		// 这样才能分文件
		if level <= zapcore.ErrorLevel {
			c := file.NewFileCore(
				file.Level(zapcore.ErrorLevel),
				file.WithName(_options.name),
				file.WithPath(_options.path))
			cores = append(cores, c)
			hooks = append(hooks, c)
		}
		if level <= zapcore.WarnLevel {
			c := file.NewFileCore(
				file.Level(zapcore.WarnLevel),
				file.WithName(_options.name),
				file.WithPath(_options.path))
			cores = append(cores, c)
			hooks = append(hooks, c)
		}
		if level <= zapcore.InfoLevel {
			c := file.NewFileCore(
				file.Level(zapcore.InfoLevel),
				file.WithName(_options.name),
				file.WithPath(_options.path))
			cores = append(cores, c)
			hooks = append(hooks, c)
		}
		if level <= zapcore.DebugLevel {
			c := file.NewFileCore(
				file.Level(zapcore.DebugLevel),
				file.WithName(_options.name),
				file.WithPath(_options.path))
			cores = append(cores, c)
			hooks = append(hooks, c)
		}
	}
	// fluent
	if len(_options.addr) > 0 {
		// encoder
		fl, err := fluent.NewFluentCore(
			fluent.Level(_options.l),
			fluent.WithName(_options.name),
			fluent.WithAddr(_options.addr),
		)
		if err != nil {
			return nil, err
		}
		cores = append(cores, fl)
	}
	core := zapcore.NewTee(cores...)
	zl := &teeCore{
		opts:  _options,
		hooks: hooks,
		core:  core,
	}
	return zl, nil
}

func (c *teeCore) Close() {
	for _, hook := range c.hooks {
		hook.Close()
	}
}

// With 复制操作
func (c *teeCore) With(fields []zapcore.Field) zapcore.Core {
	core := c.core.With(fields)
	return &teeCore{
		opts:  c.opts,
		hooks: c.hooks,
		core:  core,
	}
}

func (c *teeCore) Enabled(lvl zapcore.Level) bool {
	return c.core.Enabled(lvl)
}

func (c *teeCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	return c.core.Check(ent, ce)
}

func (c *teeCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	ent.LoggerName = c.opts.name
	return c.core.Write(ent, fields)
}

func (c *teeCore) Sync() error {
	return c.core.Sync()
}
