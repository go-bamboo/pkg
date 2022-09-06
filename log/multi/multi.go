package multi

import (
	"edu/pkg/log/core"

	"go.uber.org/zap/zapcore"
)

type multiCore struct {
	hooks []core.Logger
	core  zapcore.Core
}

// NewMultiCore new a zap logger with options.
func NewMultiCore(opts ...core.Logger) (core.Logger, error) {
	hooks := make([]zapcore.Core, 0)
	for _, opt := range opts {
		hooks = append(hooks, opt)
	}
	core := zapcore.NewTee(hooks...)
	zl := &multiCore{
		hooks: opts,
		core:  core,
	}
	return zl, nil
}

func (c *multiCore) Close() {
	for _, hook := range c.hooks {
		hook.Close()
	}
}

// With 复制操作
func (c *multiCore) With(fields []zapcore.Field) zapcore.Core {
	core := c.core.With(fields)
	return &multiCore{
		hooks: c.hooks,
		core:  core,
	}
}

func (c *multiCore) Enabled(lvl zapcore.Level) bool {
	return c.core.Enabled(lvl)
}

func (c *multiCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	return c.core.Check(ent, ce)
}

func (c *multiCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	return c.core.Write(ent, fields)
}

func (c *multiCore) Sync() error {
	return c.core.Sync()
}
