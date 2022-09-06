package hc

import (
	"fmt"
	"math"
	"time"

	"edu/pkg/log/core"

	"go.uber.org/zap/zapcore"
)

type hcCore struct {
	zapcore.LevelEnabler
	enc zapcore.Encoder
}

// NewHcCore creates a Core that writes logs to a WriteSyncer.
func NewHcCore(enc zapcore.Encoder, enab zapcore.LevelEnabler) core.Logger {
	return &hcCore{
		LevelEnabler: enab,
		enc:          enc,
	}
}

func (c *hcCore) Close() {
}

func (c *hcCore) With(fields []zapcore.Field) zapcore.Core {
	//clone := c.clone()
	//return clone
	return nil
}

func (c *hcCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(ent.Level) {
		return ce.AddCore(ent, c)
	}
	return ce
}

func (c *hcCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	//buf, err := c.enc.EncodeEntry(ent, fields)
	//if err != nil {
	//	return err
	//}
	//err = c.fl.PostWithTime(ent.Level.String(), ent.Time, buf)
	//buf.Free()
	//if err != nil {
	//	return err
	//}

	data := make(map[string]string, len(fields)+4)
	data["name"] = ent.LoggerName
	data["msg"] = ent.Message
	data["caller"] = ent.Caller.String()
	data["stack"] = ent.Stack
	for i := 0; i < len(fields); i++ {
		c.addTo(data, fields[i])
	}
	//if err := c.fl.PostWithTime(ent.Level.String(), ent.Time, data); err != nil {
	//	return err
	//}

	if ent.Level > zapcore.ErrorLevel {
		// Since we may be crashing the program, sync the output. Ignore Sync
		// errors, pending a clean solution to issue #370.
		c.Sync()
	}
	return nil
}

func (c *hcCore) Sync() error {
	return nil
}

//func (c *hcCore) clone() *fluent.fluentCore {
//	return &fluent.fluentCore{
//		LevelEnabler: c.LevelEnabler,
//		enc:          c.enc.Clone(),
//	}
//}

func (c *hcCore) addTo(enc map[string]string, f zapcore.Field) {
	var err error
	switch f.Type {
	case zapcore.ArrayMarshalerType:
		//fmt.Printf("key:%v, ArrayMarshalerType -------------%v\n", f.Key, f.Interface)
		//enc[f.Key] = f.Interface.(zapcore.ArrayMarshaler)
	case zapcore.ObjectMarshalerType:
		//fmt.Printf("key:%v, ObjectMarshalerType -------------%v\n", f.Key, f.Interface)
		//err = enc.AddObject(f.Key, f.Interface.(ObjectMarshaler))
	case zapcore.InlineMarshalerType:
		//fmt.Printf("key:%v, InlineMarshalerType -------------%v\n", f.Key, f.Interface)
		//err = f.Interface.(ObjectMarshaler).MarshalLogObject(enc)
	case zapcore.BinaryType:
		enc[f.Key] = string(f.Interface.([]byte))
	case zapcore.BoolType:
		enc[f.Key] = fmt.Sprint(f.Integer == 1)
	case zapcore.ByteStringType:
		enc[f.Key] = string(f.Interface.([]byte))
	case zapcore.Complex128Type:
		enc[f.Key] = fmt.Sprint(f.Interface.(complex128))
	case zapcore.Complex64Type:
		enc[f.Key] = fmt.Sprint(f.Interface.(complex64))
	case zapcore.DurationType:
		enc[f.Key] = fmt.Sprintf("%fs", time.Duration(f.Integer).Seconds())
	case zapcore.Float64Type:
		enc[f.Key] = fmt.Sprintf("%f", math.Float64frombits(uint64(f.Integer)))
	case zapcore.Float32Type:
		enc[f.Key] = fmt.Sprintf("%f", math.Float32frombits(uint32(f.Integer)))
	case zapcore.Int64Type:
		enc[f.Key] = fmt.Sprintf("%d", f.Integer)
	case zapcore.Int32Type:
		enc[f.Key] = fmt.Sprintf("%d", f.Integer)
	case zapcore.Int16Type:
		enc[f.Key] = fmt.Sprintf("%d", f.Integer)
	case zapcore.Int8Type:
		enc[f.Key] = fmt.Sprintf("%d", f.Integer)
	case zapcore.StringType:
		enc[f.Key] = f.String
	case zapcore.TimeType:
		if f.Interface != nil {
			enc[f.Key] = fmt.Sprint(time.Unix(0, f.Integer).In(f.Interface.(*time.Location)))
		} else {
			// Fall back to UTC if location is nil.
			enc[f.Key] = fmt.Sprint(time.Unix(0, f.Integer))
		}
	case zapcore.TimeFullType:
		enc[f.Key] = f.Interface.(time.Time).Format("2006-01-02 15:04:05")
	case zapcore.Uint64Type:
		enc[f.Key] = fmt.Sprintf("%d", f.Integer)
	case zapcore.Uint32Type:
		enc[f.Key] = fmt.Sprintf("%d", f.Integer)
	case zapcore.Uint16Type:
		enc[f.Key] = fmt.Sprintf("%d", f.Integer)
	case zapcore.Uint8Type:
		enc[f.Key] = fmt.Sprintf("%d", f.Integer)
	case zapcore.UintptrType:
		enc[f.Key] = fmt.Sprintf("%d", f.Integer)
	case zapcore.ReflectType:
		enc[f.Key] = fmt.Sprintf("%v", f.Interface)
	case zapcore.NamespaceType:
		// fmt.Printf("key:%v, NamespaceType -------------%v\n", f.Key, f.Interface)
		//fmt.Printf("-------------%v\n", f)
		//enc.OpenNamespace(f.Key)
	case zapcore.StringerType:
		enc[f.Key] = fmt.Sprintf("%v", f.Interface)
	case zapcore.ErrorType:
		enc[f.Key] = f.Interface.(error).Error()
	case zapcore.SkipType:
		break
	default:
		panic(fmt.Sprintf("unknown field type: %v", f))
	}

	if err != nil {
		enc[fmt.Sprintf("%sError", f.Key)] = err.Error()
	}
}
