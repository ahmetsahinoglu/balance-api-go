package log

import (
	"log"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func getLogLevel(level string) zapcore.Level {
	switch levelFromConfig := strings.TrimSpace(level); {
	case strings.EqualFold(levelFromConfig, "debug"):
		return zapcore.DebugLevel
	case strings.EqualFold(levelFromConfig, "error"):
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func New(logLevel string) *zap.Logger {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.Level = zap.NewAtomicLevelAt(getLogLevel(logLevel))
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := loggerConfig.Build()
	if err != nil {
		log.Fatal(err)
	}

	return logger
}
