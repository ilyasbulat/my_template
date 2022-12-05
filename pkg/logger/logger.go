package logger

import (
	"go.uber.org/zap"
)

func GetLogger(cfg zap.Config) *zap.Logger {

	logger := zap.Must(cfg.Build())

	logger.Info("logger construction succeeded")
	return logger
}
