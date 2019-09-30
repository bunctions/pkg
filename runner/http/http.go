package main

import (
	"fmt"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

func newLogger() *zap.Logger {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.OutputPaths = []string{"stdout"}
	loggerConfig.ErrorOutputPaths = []string{"stderr"}

	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}

	return logger
}

func main() {
	logger := newLogger()
	config := &config{}
	err := envconfig.Process("", config)
	if err != nil {
		logger.Panic("Error on process environment", zap.Error(err))
		return
	}

	logger.Info("Got environment", zap.Any("env", config))

	l := &loader{
		path:   config.ExportingPath,
		symbol: config.ExportingSymbol,
	}

	if err := l.load(); err != nil {
		logger.Panic("Error on loading plugin", zap.Error(err))
		return
	}

	addr := fmt.Sprintf(":%d", config.Port)
	logger.Info("Server is starting", zap.Uint("port", config.Port))

	if err := http.ListenAndServe(addr, l.handler); err != nil {
		logger.Panic("Error on starting server", zap.Error(err))
	}
}
