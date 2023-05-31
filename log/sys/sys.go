package sys

import (
	"log/syslog"

	"github.com/go-bamboo/pkg/log/core"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type sysCore struct {
	zapcore.Core
}

// NewSysCore 标准输出core
func NewSysCore(level zapcore.Level, tag string) (core.Logger, error) {
	w, err := syslog.New(syslog.LOG_DEBUG, tag)
	if err != nil {
		return nil, err
	}
	// encoder
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), // 编码器配置
		zapcore.AddSync(w),
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= level
		}), // 日志级别
	)
	return &sysCore{core}, nil
}

func (c *sysCore) Close() {
}
