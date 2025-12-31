package file

import (
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func getFilePath(dir, prefix string) string {
	return filepath.Join(dir, prefix+".log")
}

func getWriter(path, prefix string, maxSize int, maxBackups int, maxAge int) (zapcore.WriteSyncer, *lumberjack.Logger) {
	hook := lumberjack.Logger{
		Filename:   getFilePath(path, prefix), // 日志文件路径
		MaxSize:    maxSize,                   // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: maxBackups,                // 日志文件最多保存多少个备份
		MaxAge:     maxAge,                    // 文件最多保存多少天
		Compress:   false,                     // 是否压缩
		LocalTime:  true,
	}
	writer := zapcore.AddSync(&hook)
	return writer, &hook
}

func newCore(_options *options, lvl zapcore.Level) (zapcore.Core, *lumberjack.Logger) {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	w, hook := getWriter(_options.path, _options.name)
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 编码器配置
		w,
		zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return l >= lvl
		}), // 日志级别
	), hook
}
