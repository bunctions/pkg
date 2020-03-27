package http

import (
	"fmt"
	"net/http"

	"github.com/bunctions/pkg/function"
	"github.com/bunctions/pkg/runner/http/adapter"
	"github.com/bunctions/pkg/runner/util"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

const Name = "http"

func Runner() runner {
	return runner{}
}

type runner struct{}

// Start starts a HTTP runner
func (runner) Start() {
	logger := util.NewLogger().
		With(zap.String("runner", "http"))

	conf := &config{}
	err := envconfig.Process("http", conf)
	if err != nil {
		logger.Panic("Error on process environment", zap.Error(err))
		return
	}

	router := newPathRouter(function.DefaultRegistry, logger)
	handler := adapter.ApplyAll(
		router,
		adapter.NewContentTypeAdapter(conf.ContentType),
		adapter.NewLoggingAdapter(),
		adapter.NewLoggerAdapter(logger),
	)

	addr := fmt.Sprintf(":%d", conf.Port)
	logger.Info("server is starting", zap.Uint("port", conf.Port))

	if err := http.ListenAndServe(addr, handler); err != nil {
		logger.Panic("Error on starting server", zap.Error(err))
	}
}
