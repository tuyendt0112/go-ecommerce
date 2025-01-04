package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Info("hello jeztum")

	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "jezz"), zap.Int("age", 40))

	// // development
	// logger2, _ := zap.NewDevelopment()
	// logger2.Info("development call")
	// //production
	// logger3, _ := zap.NewProduction()
	// logger3.Info("production call")

	// customize
	encoder := GetEncoderLog()
	sync := GetWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)

	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("age", 3))
}

// format
func GetEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewDevelopmentEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "Time"
	encodeConfig.CallerKey = "Caller Func"
	encodeConfig.LevelKey = "Level"
	encodeConfig.MessageKey = "Message"
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}

func GetWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
