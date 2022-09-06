package log

import (
	"net/http"
	"os"
	"time"

	"edu/pkg/log/tee"

	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// DefaultLogger is default logger.
var DefaultLogger = NewLogger(tee.NewStdCore(zapcore.DebugLevel), 1)

type ZapLogger struct {
	// zap logger
	logger  *zap.Logger
	slogger *zap.SugaredLogger
}

func NewLogger(core zapcore.Core, n int) *ZapLogger {
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	skip := zap.AddCallerSkip(n)

	// 构造日志
	logger := zap.New(core, caller, skip)
	slogger := logger.Sugar()

	// copy
	return &ZapLogger{
		logger:  logger,
		slogger: slogger,
	}
}

func NewZapLogger(core zapcore.Core, n int) *ZapLogger {
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	skip := zap.AddCallerSkip(n)

	// 构造日志
	logger := zap.New(core, caller, skip)
	slogger := logger.Sugar()

	// copy
	return &ZapLogger{
		logger:  logger,
		slogger: slogger,
	}
}

// Log print the kv pairs log.
func (s *ZapLogger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 {
		return nil
	}
	if len(keyvals)%2 != 0 {
		keyvals = append(keyvals, "")
	}
	var (
		msg           string
		keysAndValues []interface{}
	)
	for i := 0; i < len(keyvals); i += 2 {
		key := keyvals[i]
		if key == "msg" {
			msg = keyvals[i+1].(string)
		} else {
			keysAndValues = append(keysAndValues, keyvals[i], keyvals[i+1])
		}
	}
	if level == log.LevelDebug {
		s.slogger.Debugw(msg, keysAndValues...)
	} else if level == log.LevelInfo {
		s.slogger.Infow(msg, keysAndValues...)
	} else if level == log.LevelWarn {
		s.slogger.Warnw(msg, keysAndValues...)
	} else if level == log.LevelError {
		s.slogger.Errorw(msg, keysAndValues...)
	}
	return nil
}

func (s *ZapLogger) Printf(format string, args ...interface{}) {
	s.slogger.Debugf(format, args...)
}

func (s *ZapLogger) Println(a ...interface{}) {
	s.slogger.Debug(a...)
}

// Debug logs a message at debug level.
func (s *ZapLogger) Debug(a ...interface{}) {
	s.slogger.Debug(a...)
}

// Debugf logs a message at debug level.
func (s *ZapLogger) Debugf(format string, a ...interface{}) {
	s.slogger.Debugf(format, a...)
}

// Debugw logs a message at debug level.
func (s *ZapLogger) Debugw(msg string, keyvals ...interface{}) {
	s.slogger.Debugw(msg, keyvals...)
}

// Info logs a message at info level.
func (s *ZapLogger) Info(a ...interface{}) {
	s.slogger.Info(a...)
}

// Infof logs a message at info level.
func (s *ZapLogger) Infof(format string, a ...interface{}) {
	s.slogger.Infof(format, a...)
}

// Infow logs a message at info level.
func (s *ZapLogger) Infow(msg string, keyvals ...interface{}) {
	s.slogger.Infow(msg, keyvals...)
}

// Warn logs a message at warn level.
func (s *ZapLogger) Warn(a ...interface{}) {
	s.slogger.Warn(a)
}

// Warnf logs a message at warnf level.
func (s *ZapLogger) Warnf(format string, a ...interface{}) {
	s.slogger.Warnf(format, a...)
}

// Warnw logs a message at warnf level.
func (s *ZapLogger) Warnw(msg string, keyvals ...interface{}) {
	s.slogger.Warnw(msg, keyvals...)
}

// Error logs a message at error level.
func (s *ZapLogger) Error(a ...interface{}) {
	s.slogger.Error(a...)
}

// Errorf logs a message at error level.
func (s *ZapLogger) Errorf(format string, a ...interface{}) {
	s.slogger.Errorf(format, a...)
}

// Errorw logs a message at error level.
func (s *ZapLogger) Errorw(msg string, keyvals ...interface{}) {
	s.slogger.Errorw(msg, keyvals...)
}

// Fatal logs a message at fatal level.
func (s *ZapLogger) Fatal(a ...interface{}) {
	s.slogger.Fatal(a...)
	os.Exit(1)
}

// Fatalf logs a message at fatal level.
func (s *ZapLogger) Fatalf(format string, a ...interface{}) {
	s.slogger.Fatalf(format, a...)
	os.Exit(1)
}

// Fatalw logs a message at fatal level.
func (s *ZapLogger) Fatalw(msg string, keyvals ...interface{}) {
	s.slogger.Fatalw(msg, keyvals...)
	os.Exit(1)
}

// LogRoundTrip prints the information about request and response.
//
func (s *ZapLogger) LogRoundTrip(
	req *http.Request,
	res *http.Response,
	err error,
	start time.Time,
	dur time.Duration,
) error {
	//var b []zapcore.Field
	//bsize := 200
	//var b = bytes.NewBuffer(make([]byte, 0, bsize))
	//var v = make([]byte, 0, bsize)
	//
	//appendTime := func(t time.Time) {
	//	v = v[:0]
	//	v = t.AppendFormat(v, time.RFC3339)
	//	b.Write(v)
	//}
	//
	//appendQuote := func(s string) {
	//	v = v[:0]
	//	v = strconv.AppendQuote(v, s)
	//	b.Write(v)
	//}
	//
	//appendInt := func(i int64) {
	//	v = v[:0]
	//	v = strconv.AppendInt(v, i, 10)
	//	b.Write(v)
	//}
	//
	//port := req.URL.Port()
	//
	//b.WriteRune('{')
	//// -- Timestamp
	//b = append(b, zapcore.Field{Key:"timestamp", Type: zapcore.TimeType, Interface: start.UTC() })
	//
	//// -- Event
	//b = append(b, zapcore.Field{Key:"event-duration", Type: zapcore.Int64Type, Integer: dur.Nanoseconds() })
	//
	//// -- URL
	//b = append(b, zapcore.Field{Key:"url-scheme", Type: zapcore.StringType, String: req.URL.Scheme })
	//b = append(b, zapcore.Field{Key:"url-domain", Type: zapcore.StringType, String: req.URL.Hostname() })
	//if port != "" {
	//	b = append(b, zapcore.Field{Key:"url-port", Type: zapcore.StringType, String: port })
	//}
	//b = append(b, zapcore.Field{Key:"url-path", Type: zapcore.StringType, String: req.URL.Hostname() })
	//b = append(b, zapcore.Field{Key:"url-query", Type: zapcore.StringType, String: req.URL.RawQuery })
	//
	//// -- HTTP
	//b = append(b, zapcore.Field{Key:"http-request-method", Type: zapcore.StringType, String: req.Method })
	//b.WriteString(`,"http":`)
	//// ---- Request
	//b.WriteString(`{"request":{`)
	//b.WriteString(`"method":`)
	//appendQuote(req.Method)
	//if l.RequestBodyEnabled() && req != nil && req.Body != nil && req.Body != http.NoBody {
	//	var buf bytes.Buffer
	//	if req.GetBody != nil {
	//		b, _ := req.GetBody()
	//		buf.ReadFrom(b)
	//	} else {
	//		buf.ReadFrom(req.Body)
	//	}
	//
	//	b.Grow(buf.Len() + 8)
	//	b.WriteString(`,"body":`)
	//	appendQuote(buf.String())
	//}
	//b.WriteRune('}') // Close "http.request"
	//// ---- Response
	//b.WriteString(`,"response":{`)
	//b.WriteString(`"status_code":`)
	//appendInt(int64(resStatusCode(res)))
	//if l.ResponseBodyEnabled() && res != nil && res.Body != nil && res.Body != http.NoBody {
	//	defer res.Body.Close()
	//	var buf bytes.Buffer
	//	buf.ReadFrom(res.Body)
	//
	//	b.Grow(buf.Len() + 8)
	//	b.WriteString(`,"body":`)
	//	appendQuote(buf.String())
	//}
	//b.WriteRune('}') // Close "http.response"
	//b.WriteRune('}') // Close "http"
	//// -- Error
	//if err != nil {
	//	b.WriteString(`,"error":{"message":`)
	//	appendQuote(err.Error())
	//	b.WriteRune('}') // Close "error"
	//}
	//b.WriteRune('}')
	//b.WriteRune('\n')
	//b.WriteTo(l.Output)

	return nil
}

// RequestBodyEnabled makes the client pass request body to logger
func (s *ZapLogger) RequestBodyEnabled() bool { return true }

// RequestBodyEnabled makes the client pass response body to logger
func (s *ZapLogger) ResponseBodyEnabled() bool { return true }
