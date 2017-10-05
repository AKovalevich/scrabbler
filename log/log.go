package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	loger *zap.Logger
)

func Initialize(logLevel string) {
	var realInfoLevel zapcore.Level

	switch logLevel {
	case zapcore.DebugLevel.String():
		realInfoLevel = zapcore.DebugLevel
		break
	case zapcore.InfoLevel.String():
		realInfoLevel = zapcore.InfoLevel
		break
	case zapcore.WarnLevel.String():
		realInfoLevel = zapcore.WarnLevel
		break
	case zapcore.ErrorLevel.String():
		realInfoLevel = zapcore.ErrorLevel
		break
	case zapcore.DPanicLevel.String():
		realInfoLevel = zapcore.DPanicLevel
		break
	case zapcore.PanicLevel.String():
		realInfoLevel = zapcore.PanicLevel
		break
	case zapcore.FatalLevel.String():
		realInfoLevel = zapcore.FatalLevel
		break
	default:
		realInfoLevel = zapcore.WarnLevel
	}

	encoderConfig := zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	config := zap.Config{
			Level:       zap.NewAtomicLevelAt(realInfoLevel),
			Development: false,
			Sampling: &zap.SamplingConfig{
				Initial:    100,
				Thereafter: 100,
			},
			Encoding:         "console",
			EncoderConfig:    encoderConfig,
			OutputPaths:      []string{"stderr"},
			ErrorOutputPaths: []string{"stderr"},
	}

	loger, _ = config.Build()
	defer loger.Sync()
}

func Info(msg string, fields ...zapcore.Field) {
	loger.Info(msg, fields...)
}

func Debug(msg string, fields ...zapcore.Field) {
	loger.Debug(msg, fields...)
}

func Warn(msg string, fields ...zapcore.Field) {
	loger.Warn(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	loger.Error(msg, fields...)
}

func Panic(msg string, fields ...zapcore.Field) {
	loger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zapcore.Field) {
	loger.Fatal(msg, fields...)
}

func String(key string, value string) zapcore.Field {
	return zap.String(key, value)
}

func Int(key string, val int) zapcore.Field {
	return zap.Int(key, val)
}

func Bool(key string, val bool) zapcore.Field {
	return zap.Bool(key, val)
}
