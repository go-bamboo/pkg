package fluent

import (
	"fmt"
	"math"
	"net"
	"net/url"
	"strconv"
	"time"

	"edu/pkg/log/core"

	"github.com/fluent/fluent-logger-golang/fluent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Option func(*options)

type options struct {
	l    zapcore.Level
	name string

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

func addFields(enc zapcore.ObjectEncoder, fields []zapcore.Field) {
	for i := range fields {
		fields[i].AddTo(enc)
	}
}

type fluentCore struct {
	opts options
	enc  zapcore.Encoder
	fl   *fluent.Fluent
}

// NewFluentCore creates a Core that writes logs to a WriteSyncer.
func NewFluentCore(opts ...Option) (core.Logger, error) {
	_options := options{
		l:    zapcore.DebugLevel,
		name: "default",
	}
	for _, o := range opts {
		o(&_options)
	}
	u, err := url.Parse(_options.addr)
	if err != nil {
		return nil, err
	}
	c := fluent.Config{
		Timeout:            _options.timeout,
		WriteTimeout:       _options.writeTimeout,
		BufferLimit:        _options.bufferLimit,
		RetryWait:          _options.retryWait,
		MaxRetry:           _options.maxRetry,
		MaxRetryWait:       _options.maxRetryWait,
		TagPrefix:          _options.tagPrefix,
		Async:              _options.async,
		ForceStopAsyncSend: _options.forceStopAsyncSend,
	}
	switch u.Scheme {
	case "tcp":
		host, port, err2 := net.SplitHostPort(u.Host)
		if err2 != nil {
			return nil, err2
		}
		if c.FluentPort, err = strconv.Atoi(port); err != nil {
			return nil, err
		}
		c.FluentNetwork = u.Scheme
		c.FluentHost = host
	case "unix":
		c.FluentNetwork = u.Scheme
		c.FluentSocketPath = u.Path
	default:
		return nil, fmt.Errorf("unknown network: %s", u.Scheme)
	}
	fl, err := fluent.New(c)
	if err != nil {
		return nil, err
	}
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	return &fluentCore{
		opts: _options,
		enc:  zapcore.NewJSONEncoder(encoderConfig),
		fl:   fl,
	}, nil
}

func (c *fluentCore) Close() {
	c.fl.Close()
}

func (c *fluentCore) With(fields []zapcore.Field) zapcore.Core {
	clone := c.clone()
	addFields(clone.enc, fields)
	return clone
}

func (c *fluentCore) Enabled(lvl zapcore.Level) bool {
	return lvl >= c.opts.l
}

func (c *fluentCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if ent.Level >= c.opts.l {
		return ce.AddCore(ent, c)
	}
	return ce
}

func (c *fluentCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	data := make(map[string]string, len(fields)+4)
	data["name"] = ent.LoggerName
	data["msg"] = ent.Message
	data["caller"] = ent.Caller.String()
	data["stack"] = ent.Stack
	for i := 0; i < len(fields); i++ {
		c.addTo(data, fields[i])
	}
	if err := c.fl.PostWithTime(ent.Level.String(), ent.Time, data); err != nil {
		return err
	}

	if ent.Level > zapcore.ErrorLevel {
		// Since we may be crashing the program, sync the output. Ignore Sync
		// errors, pending a clean solution to issue #370.
		c.Sync()
	}
	return nil
}

func (c *fluentCore) Sync() error {
	return nil
}

func (c *fluentCore) clone() *fluentCore {
	return &fluentCore{
		enc: c.enc.Clone(),
		fl:  c.fl,
	}
}

func (c *fluentCore) addTo(enc map[string]string, f zapcore.Field) {
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
