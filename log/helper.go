package log

import (
	"github.com/go-bamboo/pkg/log/core"
	"github.com/go-bamboo/pkg/log/multi"
	"go.uber.org/zap/zapcore"
)

func WithCore(c zapcore.Core, kv ...interface{}) zapcore.Core {
	if len(kv) == 0 {
		return c
	}
	if len(kv)%2 != 0 {
		kv = append(kv, "")
	}
	var keysAndValues []zapcore.Field
	for i := 0; i < len(kv); i += 2 {
		key := kv[i]
		val := kv[i+1]
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

func Init(c []*core.Conf, opts ...Option) core.Logger {
	hooks := make([]core.Logger, 0)
	for _, conf := range c {
		co, err := core.Create(conf)
		if err != nil {
			Fatal(err)
		}
		hooks = append(hooks, co)
	}
	co, err := multi.NewMultiCore(hooks...)
	if err != nil {
		panic(err)
	}

	// global
	logger := NewLogger(co, opts...)
	SetLogger(logger)
	return co
}
