package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GetLogger(cfg zap.Config) *zap.Logger {
	cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	logger := zap.Must(cfg.Build(zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel)))
	logger.Info("logger construction succeeded")
	return logger
}
