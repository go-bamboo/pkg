package file

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"path/filepath"
)

func getFilePath(dir, prefix string, lvl zapcore.Level) string {
	return filepath.Join(dir, prefix+"-"+lvl.String()+".log")
}

func getWriter(path, prefix string, lvl zapcore.Level) (zapcore.WriteSyncer, *lumberjack.Logger) {
	hook := lumberjack.Logger{
		Filename:   getFilePath(path, prefix, lvl), // 日志文件路径
		MaxSize:    128,                            // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,                             // 日志文件最多保存多少个备份
		MaxAge:     7,                              // 文件最多保存多少天
		Compress:   false,                          // 是否压缩
	}
	writer := zapcore.AddSync(&hook)
	return writer, &hook
}

func newCore(_options *options, lvl zapcore.Level) (zapcore.Core, *lumberjack.Logger) {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	w, hook := getWriter(_options.path, _options.name, lvl)
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 编码器配置
		w,
		zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return l == lvl
		}), // 日志级别
	), hook
}
