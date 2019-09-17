package io

import (
	"context"
	"io"
	"strings"
)

type ctxInputReaderKey struct{}

func WithInputReader(ctx context.Context, r io.ReadCloser) context.Context {
	parent := ctx
	if parent == nil {
		parent = context.Background()
	}

	return context.WithValue(parent, ctxInputReaderKey{}, r)
}

func InputReaderFromContext(ctx context.Context) io.Reader {
	emptyReader := strings.NewReader("")
	if ctx == nil {
		return emptyReader
	}

	r, ok := ctx.Value(ctxInputReaderKey{}).(io.Reader)
	if !ok {
		return emptyReader
	}

	return r
}
