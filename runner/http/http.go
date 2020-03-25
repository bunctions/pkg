package http

import (
	"fmt"
	"net/http"

	"github.com/bunctions/pkg/function"
	"github.com/bunctions/pkg/runner/http/adapter"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

// Start starts a HTTP runner
func Start() {
	logger := newLogger()

	conf := &config{}
	err := envconfig.Process("http", conf)
	if err != nil {
		logger.Panic("Error on process environment", zap.Error(err))
		return
	}

	handler := adapter.ApplyAll(
		newPathRouter(function.DefaultRegistry),
		adapter.NewContentTypeAdapter(conf.ContentType),
	)

	addr := fmt.Sprintf(":%d", conf.Port)
	logger.Info("Server is starting", zap.Uint("port", conf.Port))

	if err := http.ListenAndServe(addr, handler); err != nil {
		logger.Panic("Error on starting server", zap.Error(err))
	}
}
