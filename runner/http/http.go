package main

import (
	"fmt"
	"net/http"

	"github.com/bunctions/pkg/runner/http/adapter"
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
	conf := &config{}
	err := envconfig.Process("", conf)
	if err != nil {
		logger.Panic("Error on process environment", zap.Error(err))
		return
	}

	l := &loader{
		path:   conf.ExportingPath,
		symbol: conf.ExportingSymbol,
	}

	handler, err := l.loadHandler()
	if err != nil {
		logger.Panic("Error on loading plugin", zap.Error(err))
		return
	}

	handler = adapter.ApplyAll(
		handler,
		adapter.NewContentTypeAdapter(conf.ContentType),
	)

	addr := fmt.Sprintf(":%d", conf.Port)
	logger.Info("Server is starting", zap.Uint("port", conf.Port))

	if err := http.ListenAndServe(addr, handler); err != nil {
		logger.Panic("Error on starting server", zap.Error(err))
	}
}
