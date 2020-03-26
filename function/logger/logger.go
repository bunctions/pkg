package logger

import (
	"context"

	"go.uber.org/zap"
)

type ctxLoggerKey struct{}

// ContextWithLogger returns a copy of given context, with
// the logger injected.
func ContextWithLogger(
	ctx context.Context,
	logger *zap.Logger,
) context.Context {
	parent := ctx
	if parent == nil {
		parent = context.Background()
	}

	return context.WithValue(parent, ctxLoggerKey{}, logger)
}

func LoggerFromContext(ctx context.Context) *zap.Logger {
	nopLogger := zap.NewNop()
	if ctx == nil {
		return nopLogger
	}

	if l, ok := ctx.Value(ctxLoggerKey{}).(*zap.Logger); ok {
		return l
	}

	return nopLogger
}
