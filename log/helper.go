package log

import (
	"github.com/go-bamboo/pkg/log/aws"
	"github.com/go-bamboo/pkg/log/core"
	"github.com/go-bamboo/pkg/log/file"
	"github.com/go-bamboo/pkg/log/fluent"
	"github.com/go-bamboo/pkg/log/multi"
	"github.com/go-bamboo/pkg/log/std"
	"go.uber.org/zap/zapcore"
)

func With(c zapcore.Core, kv ...interface{}) zapcore.Core {
	if len(kv) == 0 {
		return c
	}
	if len(kv)%2 != 0 {
		kv = append(kv, "")
	}
	var keysAndValues []zapcore.Field
	for i := 0; i < len(kv); i += 2 {
		key := kv[i]
		keysAndValues = append(keysAndValues, zapcore.Field{Key: key.(string), Type: zapcore.StringType, String: kv[i+1].(string)})
	}
	return c.With(keysAndValues)
}

func NewLoggerCore(c *Conf) (core.Logger, error) {
	hooks := make([]core.Logger, 0)
	if c.Console.Enable {
		c := std.NewStdCore(zapcore.Level(c.Console.Level))
		hooks = append(hooks, c)
	}
	if c.File.Enable {
		c := file.NewFileCore(
			file.Level(zapcore.Level(c.File.Level)),
			file.WithPath(c.File.Path),
			file.WithName(c.File.Name),
		)
		hooks = append(hooks, c)
	}
	if c.Fluent.Enable {
		c, err := fluent.NewFluentCore(
			fluent.Level(zapcore.Level(c.Fluent.Level)),
			fluent.WithAddr(c.Fluent.Addr),
		)
		if err != nil {
			return nil, err
		}
		hooks = append(hooks, c)
	}
	if c.CloudWatch.Enable {
		c, err := aws.NewCloudWatchCore(
			aws.Level(zapcore.Level(c.CloudWatch.Level)),
			aws.WithAccessKey(c.CloudWatch.Key),
			aws.WithAccessSecret(c.CloudWatch.Secret),
		)
		if err != nil {
			return nil, err
		}
		hooks = append(hooks, c)
	}
	logger, err := multi.NewMultiCore(hooks...)
	if err != nil {
		return nil, err
	}
	return logger, nil
}

func Init(c *Conf) core.Logger {
	core, err := NewLoggerCore(c)
	if err != nil {
		Fatal(err)
	}
	// global
	logger := NewLogger(core, 1)
	SetLogger(logger)
	return core
}
