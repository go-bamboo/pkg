package file

import (
	"github.com/go-bamboo/pkg/log/core"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Option func(*options)

type options struct {
	l    zapcore.Level
	name string
	path string
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

func WithPath(path string) Option {
	return func(c *options) {
		c.path = path
	}
}

// wrap tee
type fileCore struct {
	opts  options
	hooks []*lumberjack.Logger
	core  zapcore.Core
}

// NewFileCore new a zap logger with options.
func NewFileCore(opts ...Option) core.Logger {
	_options := options{
		l:    zapcore.DebugLevel,
		name: "default",
	}
	for _, o := range opts {
		o(&_options)
	}
	if _options.l < zapcore.DebugLevel {
		_options.l = zapcore.DebugLevel
	}
	if len(_options.name) < 0 {
		_options.name = "default"
	}
	hooks := make([]*lumberjack.Logger, 0)
	cores := make([]zapcore.Core, 0)
	if _options.l <= zapcore.ErrorLevel {
		c, hook := newCore(&_options, zapcore.ErrorLevel)
		cores = append(cores, c)
		hooks = append(hooks, hook)
	}
	if _options.l <= zapcore.WarnLevel {
		c, hook := newCore(&_options, zapcore.WarnLevel)
		cores = append(cores, c)
		hooks = append(hooks, hook)
	}
	if _options.l <= zapcore.InfoLevel {
		c, hook := newCore(&_options, zapcore.InfoLevel)
		cores = append(cores, c)
		hooks = append(hooks, hook)
	}
	if _options.l <= zapcore.DebugLevel {
		c, hook := newCore(&_options, zapcore.DebugLevel)
		cores = append(cores, c)
		hooks = append(hooks, hook)
	}
	c := zapcore.NewTee(cores...)
	return &fileCore{
		opts:  _options,
		hooks: hooks,
		core:  c,
	}
}

func (c *fileCore) Close() {
	for _, hook := range c.hooks {
		hook.Close()
	}
}

func (c *fileCore) Enabled(lvl zapcore.Level) bool {
	return c.core.Enabled(lvl)
}

// With 复制操作
func (c *fileCore) With(fields []zapcore.Field) zapcore.Core {
	core := c.core.With(fields)
	return &fileCore{
		opts:  c.opts,
		hooks: c.hooks,
		core:  core,
	}
}

func (c *fileCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	return c.core.Check(ent, ce)
}

func (c *fileCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	ent.LoggerName = c.opts.name
	return c.core.Write(ent, fields)
}

func (c *fileCore) Sync() error {
	return c.core.Sync()
}
