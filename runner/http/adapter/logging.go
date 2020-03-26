package adapter

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"

	funclogger "github.com/bunctions/pkg/function/logger"
)

func NewLoggingAdapter() Adapter {
	return AdapterFunc(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			logger := funclogger.LoggerFromContext(r.Context())
			logger = logger.With(zap.String("request_id", uuid.New().String()))

			injectedRequest := r.WithContext(
				funclogger.ContextWithLogger(r.Context(), logger),
			)

			startTime := time.Now()
			next.ServeHTTP(rw, injectedRequest)

			logger.Info(
				"handle request",
				zap.String("path", r.URL.Path),
				zap.String("method", r.Method),
				zap.Duration("duration", time.Since(startTime)),
			)
		})
	})
}
