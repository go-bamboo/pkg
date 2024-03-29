package std

import (
	"os"

	"github.com/go-bamboo/pkg/log/core"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type stdCore struct {
	zapcore.Core
}

// NewStdCore 标准输出core
func NewStdCore(level zapcore.Level) core.Logger {
	// encoder
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	c := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), // 编码器配置
		zapcore.AddSync(os.Stdout),
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= level
		}), // 日志级别
	)
	return &stdCore{c}
}

func (c *stdCore) Close() {
}
