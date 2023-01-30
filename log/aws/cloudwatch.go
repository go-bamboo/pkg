package aws

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/go-bamboo/pkg/log/core"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const queueCap = 1000

type cloudWatchCore struct {
	zapcore.LevelEnabler
	enc zapcore.Encoder
	// out
	out *cloudWatchWriteSyncer
}

type options struct {
	l            zapcore.Level
	accessKey    string
	accessSecret string
	sessionToken string
	endpoint     string
	profile      string
	region       string
	logGroupName string
}

func defaultOptions() *options {
	return &options{
		profile: "default",
		region:  "app",
	}
}

func Level(l zapcore.Level) Option {
	return func(c *options) {
		c.l = l
	}
}

func WithRegion(region string) Option {
	return func(alc *options) {
		alc.region = region
	}
}

func WithAccessKey(ak string) Option {
	return func(alc *options) {
		alc.accessKey = ak
	}
}

func WithAccessSecret(as string) Option {
	return func(alc *options) {
		alc.accessSecret = as
	}
}

func WithSessionToken(sk string) Option {
	return func(alc *options) {
		alc.sessionToken = sk
	}
}

func WithEndpoint(endpoint string) Option {
	return func(alc *options) {
		alc.endpoint = endpoint
	}
}

func WithLogGroupName(logGroupName string) Option {
	return func(alc *options) {
		alc.logGroupName = logGroupName
	}
}

type Option func(alc *options)

// NewCloudWatchCore new a cloud watch logger with options.
func NewCloudWatchCore(options ...Option) (c core.Logger, err error) {
	opts := defaultOptions()
	for _, o := range options {
		o(opts)
	}
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	logger := &cloudWatchCore{
		LevelEnabler: zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= opts.l
		}),
		enc: zapcore.NewJSONEncoder(encoderConfig),
	}
	out, err := newCloudWatchWriteSyncer(opts)
	if err != nil {
		return nil, err
	}
	logger.out = out
	return logger, nil
}

func (c *cloudWatchCore) Close() {
	c.out.Close()
}

func (c *cloudWatchCore) Level() zapcore.Level {
	return zapcore.LevelOf(c.LevelEnabler)
}

func (c *cloudWatchCore) With(fields []zapcore.Field) zapcore.Core {
	clone := c.clone()
	addFields(clone.enc, fields)
	return clone
}

func (c *cloudWatchCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(ent.Level) {
		return ce.AddCore(ent, c)
	}
	return ce
}

func (c *cloudWatchCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {

	//data := make(map[string]string, len(fields)+4)
	//data["name"] = ent.LoggerName
	//data["msg"] = ent.Message
	//data["caller"] = ent.Caller.String()
	//data["stack"] = ent.Stack
	//data["level"] = ent.Level.String()
	//for _, field := range fields {
	//	data[field.Key] = toString(field)
	//}
	//msg, err := json.Marshal(data)
	//if err != nil {
	//	return err
	//}

	buf, err := c.enc.EncodeEntry(ent, fields)
	if err != nil {
		return err
	}
	ev := new(types.InputLogEvent)
	ev.Message = aws.String(string(buf.Bytes()))
	ev.Timestamp = aws.Int64(time.Now().UnixNano() / int64(time.Millisecond))
	c.out.Write(ev)
	buf.Free()
	return nil
}

func (c *cloudWatchCore) Sync() error {
	return c.out.Sync()
}

func (c *cloudWatchCore) clone() *cloudWatchCore {
	return &cloudWatchCore{
		LevelEnabler: c.LevelEnabler,
		enc:          c.enc.Clone(),
		out:          c.out,
	}
}

func addFields(enc zapcore.ObjectEncoder, fields []zapcore.Field) {
	for i := range fields {
		fields[i].AddTo(enc)
	}
}

// toString 任意类型转string
//func toString(f zapcore.Field) string {
//	var key string
//	switch f.Type {
//	case zapcore.BinaryType:
//		key = string(f.Interface.([]byte))
//	case zapcore.BoolType:
//		key = fmt.Sprint(f.Integer == 1)
//	case zapcore.ByteStringType:
//		key = string(f.Interface.([]byte))
//	case zapcore.Complex128Type:
//		key = fmt.Sprint(f.Interface.(complex128))
//	case zapcore.Complex64Type:
//		key = fmt.Sprint(f.Interface.(complex64))
//	case zapcore.DurationType:
//		key = fmt.Sprintf("%fs", time.Duration(f.Integer).Seconds())
//	case zapcore.Float64Type:
//		key = fmt.Sprintf("%f", math.Float64frombits(uint64(f.Integer)))
//	case zapcore.Float32Type:
//		key = fmt.Sprintf("%f", math.Float32frombits(uint32(f.Integer)))
//	case zapcore.Int64Type:
//		key = fmt.Sprintf("%d", f.Integer)
//	case zapcore.Int32Type:
//		key = fmt.Sprintf("%d", f.Integer)
//	case zapcore.Int16Type:
//		key = fmt.Sprintf("%d", f.Integer)
//	case zapcore.Int8Type:
//		key = fmt.Sprintf("%d", f.Integer)
//	case zapcore.StringType:
//		key = f.String
//	case zapcore.TimeType:
//		if f.Interface != nil {
//			key = fmt.Sprint(time.Unix(0, f.Integer).In(f.Interface.(*time.Location)))
//		} else {
//			// Fall back to UTC if location is nil.
//			key = fmt.Sprint(time.Unix(0, f.Integer))
//		}
//	case zapcore.TimeFullType:
//		key = f.Interface.(time.Time).Format("2006-01-02 15:04:05")
//	case zapcore.Uint64Type:
//		key = fmt.Sprintf("%d", f.Integer)
//	case zapcore.Uint32Type:
//		key = fmt.Sprintf("%d", f.Integer)
//	case zapcore.Uint16Type:
//		key = fmt.Sprintf("%d", f.Integer)
//	case zapcore.Uint8Type:
//		key = fmt.Sprintf("%d", f.Integer)
//	case zapcore.UintptrType:
//		key = fmt.Sprintf("%d", f.Integer)
//	case zapcore.ReflectType:
//		key = fmt.Sprintf("%v", f.Interface)
//	}
//	return key
//}
