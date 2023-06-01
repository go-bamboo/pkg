package gormx

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/logger"
)

// Colors
const (
	Reset       = "\033[0m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Blue        = "\033[34m"
	Magenta     = "\033[35m"
	Cyan        = "\033[36m"
	White       = "\033[37m"
	BlueBold    = "\033[34;1m"
	MagentaBold = "\033[35;1m"
	RedBold     = "\033[31;1m"
	YellowBold  = "\033[33;1m"
)

type Logger struct {
	c                                   *LoggerConf
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string

	level   logger.LogLevel
	logger  *zap.Logger
	slogger *zap.SugaredLogger
}

func NewLogger(config *LoggerConf, core zapcore.Core) logger.Interface {
	// gorm
	var (
		infoStr      = "%s\n"
		warnStr      = "%s\n"
		errStr       = "%s\n"
		traceStr     = "%s\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	)

	if config.Colorful {
		infoStr = Green + "%s\n" + Reset + Green
		warnStr = BlueBold + "%s\n" + Reset + Magenta
		errStr = Magenta + "%s\n" + Reset + Red
		traceStr = Reset + Yellow + "[%.3fms] " + BlueBold + "[rows:%v]" + Reset + " %s"
		traceWarnStr = Green + "%s " + Reset + RedBold + "[%.3fms] " + Yellow + "[rows:%v]" + Magenta + " %s" + Reset
		traceErrStr = RedBold + "%s " + Reset + Yellow + "[%.3fms] " + BlueBold + "[rows:%v]" + Reset + " %s"
	}

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	skip := zap.AddCallerSkip(1)
	//level := zap.IncreaseLevel(zapcore.LevelOf(config.LogLevel))

	// 构造日志
	zapLogger := zap.New(core, caller, skip)
	zapSugarLogger := zapLogger.Sugar()

	l := &Logger{
		c: config,
		// gorm
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,

		// gorm log
		level: logger.LogLevel(config.LogLevel),

		// zap
		logger:  zapLogger,
		slogger: zapSugarLogger,
	}
	return l
}

// LogMode log mode
func (l *Logger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.level = level
	return &newLogger
}

// Info print info
func (l *Logger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= logger.Info {
		l.slogger.Infof(l.infoStr+msg, data...)
	}
}

// Warn print warn messages
func (l *Logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= logger.Warn {
		l.slogger.Warnf(l.warnStr+msg, data...)
	}
}

// Error print error messages
func (l *Logger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= logger.Error {
		l.slogger.Errorf(l.errStr+msg, data...)
	}
}

// Trace print sql message
func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.level <= logger.Silent {
		return
	}
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.level >= logger.Error && (!IsGormErrRecordNotFound(err) || !l.c.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			l.slogger.Errorf(l.traceErrStr, err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.slogger.Errorf(l.traceErrStr, err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.c.SlowThreshold.AsDuration() && l.c.SlowThreshold.AsDuration() != 0 && l.level >= logger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.c.SlowThreshold)
		if rows == -1 {
			l.slogger.Warnf(l.traceWarnStr, slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.slogger.Warnf(l.traceWarnStr, slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case l.level == logger.Info:
		sql, rows := fc()
		if rows == -1 {
			l.slogger.Infof(l.traceStr, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.slogger.Infof(l.traceStr, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
