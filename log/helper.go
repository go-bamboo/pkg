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
		val := kv[2]
		switch val.(type) {
		case int:
			keysAndValues = append(keysAndValues, zapcore.Field{Key: key.(string), Type: zapcore.Int32Type, Integer: int64(val.(int))})
		case uint:
			keysAndValues = append(keysAndValues, zapcore.Field{Key: key.(string), Type: zapcore.Uint32Type, Integer: int64(val.(uint))})
		case int8:
			keysAndValues = append(keysAndValues, zapcore.Field{Key: key.(string), Type: zapcore.Int32Type, Integer: int64(val.(int8))})
		case uint8:
			keysAndValues = append(keysAndValues, zapcore.Field{Key: key.(string), Type: zapcore.Uint32Type, Integer: int64(val.(uint8))})
		case int16:
			keysAndValues = append(keysAndValues, zapcore.Field{Key: key.(string), Type: zapcore.Int16Type, Integer: int64(val.(int16))})
		case uint16:
			keysAndValues = append(keysAndValues, zapcore.Field{Key: key.(string), Type: zapcore.Uint16Type, Integer: int64(val.(uint16))})
		case int32:
			keysAndValues = append(keysAndValues, zapcore.Field{Key: key.(string), Type: zapcore.Int32Type, Integer: int64(val.(int32))})
		case uint32:
			keysAndValues = append(keysAndValues, zapcore.Field{Key: key.(string), Type: zapcore.Uint32Type, Integer: int64(val.(uint32))})
		case int64:
			keysAndValues = append(keysAndValues, zapcore.Field{Key: key.(string), Type: zapcore.Int64Type, Integer: int64(val.(int64))})
		case uint64:
			keysAndValues = append(keysAndValues, zapcore.Field{Key: key.(string), Type: zapcore.Uint64Type, Integer: int64(val.(uint64))})
		case string:
			keysAndValues = append(keysAndValues, zapcore.Field{Key: key.(string), Type: zapcore.StringType, String: val.(string)})
		}
	}
	return c.With(keysAndValues)
}

func NewLoggerCore(c *Conf) (core.Logger, error) {
	hooks := make([]core.Logger, 0)
	if c.Console != nil && c.Console.Enable {
		c := std.NewStdCore(zapcore.Level(c.Console.Level))
		hooks = append(hooks, c)
	}
	if c.File != nil && c.File.Enable {
		c := file.NewFileCore(
			file.Level(zapcore.Level(c.File.Level)),
			file.WithPath(c.File.Path),
			file.WithName(c.File.Name),
		)
		hooks = append(hooks, c)
	}
	if c.Fluent != nil && c.Fluent.Enable {
		c, err := fluent.NewFluentCore(
			fluent.Level(zapcore.Level(c.Fluent.Level)),
			fluent.WithAddr(c.Fluent.Addr),
		)
		if err != nil {
			return nil, err
		}
		hooks = append(hooks, c)
	}
	if c.CloudWatch != nil && c.CloudWatch.Enable {
		c, err := aws.NewCloudWatchCore(
			aws.Level(zapcore.Level(c.CloudWatch.Level)),
			aws.WithRegion(c.CloudWatch.Region),
			aws.WithAccessKey(c.CloudWatch.Key),
			aws.WithAccessSecret(c.CloudWatch.Secret),
			aws.WithLogGroupName(c.CloudWatch.LogGroupName),
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
