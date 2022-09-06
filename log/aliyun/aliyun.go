package aliyun

import (
	"fmt"
	"math"
	"time"

	"edu/pkg/log/core"

	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"go.uber.org/zap/zapcore"
	"google.golang.org/protobuf/proto"
)

type aliyunCore struct {
	producer *producer.Producer
	opts     *options
}

type options struct {
	l            zapcore.Level
	accessKey    string
	accessSecret string
	endpoint     string
	project      string
	logstore     string
}

func defaultOptions() *options {
	return &options{
		project:  "projectName",
		logstore: "app",
	}
}

func WithEndpoint(endpoint string) Option {
	return func(alc *options) {
		alc.endpoint = endpoint
	}
}

func WithProject(project string) Option {
	return func(alc *options) {
		alc.project = project
	}
}

func WithLogstore(logstore string) Option {
	return func(alc *options) {
		alc.logstore = logstore
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

type Option func(alc *options)

// NewAliyunLog new a aliyun logger with options.
func NewAliyunCore(options ...Option) core.Logger {
	opts := defaultOptions()
	for _, o := range options {
		o(opts)
	}

	producerConfig := producer.GetDefaultProducerConfig()
	producerConfig.Endpoint = opts.endpoint
	producerConfig.AccessKeyID = opts.accessKey
	producerConfig.AccessKeySecret = opts.accessSecret
	producerInst := producer.InitProducer(producerConfig)

	return &aliyunCore{
		opts:     opts,
		producer: producerInst,
	}
}

func (c *aliyunCore) Close() {
	timeoutMs := time.Second * 1
	c.producer.Close(timeoutMs.Microseconds())
}

func (c *aliyunCore) With(fields []zapcore.Field) zapcore.Core {
	//clone := c.clone()
	//addFields(clone.enc, fields)
	//return clone
	return nil
}

func (c *aliyunCore) Enabled(lvl zapcore.Level) bool {
	return lvl >= c.opts.l
}

func (c *aliyunCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if ent.Level >= c.opts.l {
		return ce.AddCore(ent, c)
	}
	return ce
}

func (c *aliyunCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	buf := ent.Level.String()
	levelTitle := "level"

	// level
	contents := make([]*sls.LogContent, len(fields)+4)
	contents = append(contents, &sls.LogContent{
		Key:   &levelTitle,
		Value: &buf,
	})

	for i := 0; i < len(fields); i++ {
		key := fields[i].Key
		value := toString(&fields[i])
		contents = append(contents, &sls.LogContent{
			Key:   &key,
			Value: &value,
		})
	}

	logInst := &sls.Log{
		Time:     proto.Uint32(uint32(time.Now().Unix())),
		Contents: contents,
	}

	err := c.producer.SendLog(c.opts.project, c.opts.logstore, "", "", logInst)

	return err
}

func (c *aliyunCore) Sync() error {
	return nil
}

// toString 任意类型转string
func toString(f *zapcore.Field) string {
	var key string
	if f == nil {
		return key
	}
	switch f.Type {
	case zapcore.BinaryType:
		key = string(f.Interface.([]byte))
	case zapcore.BoolType:
		key = fmt.Sprint(f.Integer == 1)
	case zapcore.ByteStringType:
		key = string(f.Interface.([]byte))
	case zapcore.Complex128Type:
		key = fmt.Sprint(f.Interface.(complex128))
	case zapcore.Complex64Type:
		key = fmt.Sprint(f.Interface.(complex64))
	case zapcore.DurationType:
		key = fmt.Sprintf("%fs", time.Duration(f.Integer).Seconds())
	case zapcore.Float64Type:
		key = fmt.Sprintf("%f", math.Float64frombits(uint64(f.Integer)))
	case zapcore.Float32Type:
		key = fmt.Sprintf("%f", math.Float32frombits(uint32(f.Integer)))
	case zapcore.Int64Type:
		key = fmt.Sprintf("%d", f.Integer)
	case zapcore.Int32Type:
		key = fmt.Sprintf("%d", f.Integer)
	case zapcore.Int16Type:
		key = fmt.Sprintf("%d", f.Integer)
	case zapcore.Int8Type:
		key = fmt.Sprintf("%d", f.Integer)
	case zapcore.StringType:
		key = f.String
	case zapcore.TimeType:
		if f.Interface != nil {
			key = fmt.Sprint(time.Unix(0, f.Integer).In(f.Interface.(*time.Location)))
		} else {
			// Fall back to UTC if location is nil.
			key = fmt.Sprint(time.Unix(0, f.Integer))
		}
	case zapcore.TimeFullType:
		key = f.Interface.(time.Time).Format("2006-01-02 15:04:05")
	case zapcore.Uint64Type:
		key = fmt.Sprintf("%d", f.Integer)
	case zapcore.Uint32Type:
		key = fmt.Sprintf("%d", f.Integer)
	case zapcore.Uint16Type:
		key = fmt.Sprintf("%d", f.Integer)
	case zapcore.Uint8Type:
		key = fmt.Sprintf("%d", f.Integer)
	case zapcore.UintptrType:
		key = fmt.Sprintf("%d", f.Integer)
	case zapcore.ReflectType:
		key = fmt.Sprintf("%v", f.Interface)
	}
	return key
}
