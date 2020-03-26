package util

import "go.uber.org/zap"

func NewLogger() *zap.Logger {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.OutputPaths = []string{"stdout"}
	loggerConfig.ErrorOutputPaths = []string{"stderr"}

	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}

	return logger
}
