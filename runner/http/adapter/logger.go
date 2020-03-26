package adapter

import (
	"net/http"

	"github.com/bunctions/pkg/function/logger"
	"go.uber.org/zap"
)

type loggerAdapter struct {
	logger *zap.Logger
}

func NewLoggerAdapter(logger *zap.Logger) Adapter {
	return &loggerAdapter{
		logger: logger,
	}
}

func (a *loggerAdapter) Apply(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		injectedRequest := r.WithContext(
			logger.ContextWithLogger(r.Context(), a.logger),
		)

		next.ServeHTTP(rw, injectedRequest)
	})
}
