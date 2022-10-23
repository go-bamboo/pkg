package file

import (
	"path/filepath"

	"github.com/go-bamboo/pkg/log/core"
	"go.uber.org/zap"
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

func getFilePath(dir, prefix string, lvl zapcore.Level) string {
	if lvl == zapcore.DebugLevel {
		return filepath.Join(dir, prefix+"-debug") + ".log"
	} else if lvl == zapcore.InfoLevel {
		return filepath.Join(dir, prefix+"-info") + ".log"
	} else if lvl == zapcore.WarnLevel {
		return filepath.Join(dir, prefix+"-warn") + ".log"
	} else {
		return filepath.Join(dir, prefix+"-error") + ".log"
	}
}

func getWriter(path, prefix string, lvl zapcore.Level) (zapcore.WriteSyncer, *lumberjack.Logger) {
	hook := lumberjack.Logger{
		Filename:   getFilePath(path, prefix, lvl), // 日志文件路径
		MaxSize:    128,                            // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,                             // 日志文件最多保存多少个备份
		MaxAge:     7,                              // 文件最多保存多少天
		Compress:   false,                          // 是否压缩
	}
	writer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook))
	return writer, &hook
}

type fileCore struct {
	opts options

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
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	hooks := make([]*lumberjack.Logger, 0)
	cores := make([]zapcore.Core, 0)
	if _options.l <= zapcore.ErrorLevel {
		w, hook := getWriter(_options.path, _options.name, zapcore.ErrorLevel)
		cores = append(cores, zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig), // 编码器配置
			w,
			zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl == zapcore.ErrorLevel
			}), // 日志级别
		))
		hooks = append(hooks, hook)
	}
	if _options.l <= zapcore.WarnLevel {
		w, hook := getWriter(_options.path, _options.name, zapcore.WarnLevel)
		cores = append(cores, zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig), // 编码器配置
			w,
			zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl == zapcore.WarnLevel
			}), // 日志级别
		))
		hooks = append(hooks, hook)
	}
	if _options.l <= zapcore.InfoLevel {
		w, hook := getWriter(_options.path, _options.name, zapcore.InfoLevel)
		cores = append(cores, zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig), // 编码器配置
			w,
			zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl == zapcore.InfoLevel
			}), // 日志级别
		))
		hooks = append(hooks, hook)
	}
	if _options.l <= zapcore.DebugLevel {
		w, hook := getWriter(_options.path, _options.name, zapcore.DebugLevel)
		cores = append(cores, zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig), // 编码器配置
			w,
			zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl == zapcore.DebugLevel
			}), // 日志级别
		))
		hooks = append(hooks, hook)
	}
	core := zapcore.NewTee(cores...)
	return &fileCore{
		opts:  _options,
		hooks: hooks,
		core:  core,
	}
}

func (c *fileCore) Close() {
	for _, hook := range c.hooks {
		hook.Close()
	}
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

func (c *fileCore) Enabled(lvl zapcore.Level) bool {
	return lvl >= c.opts.l
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
