package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *zap.SugaredLogger

var logPath = "./log/webalbums.log"

func Initialize() {
	consoleCore := createConsoleCore()
	fileCore := createFileCore(logPath)
	zapLog := zap.New(zapcore.NewTee(fileCore, consoleCore), zap.AddCaller())
	zap.ReplaceGlobals(zapLog)

	Log = zap.S()
}

// 创建File Core
func createFileCore(filePath string) zapcore.Core {
	syncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    500,
		MaxAge:     30,
		MaxBackups: 30,
		LocalTime:  true,
		Compress:   false,
	})
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zapcore.DebugLevel)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
		// FunctionKey:    "F",
	}
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	return zapcore.NewCore(encoder, syncer, atomicLevel)
}

func createConsoleCore() zapcore.Core {
	consoleWriter := zapcore.Lock(os.Stdout)
	consoleEncoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
		// FunctionKey: "F",
	})
	return zapcore.NewCore(consoleEncoder, consoleWriter, zapcore.DebugLevel)
}
