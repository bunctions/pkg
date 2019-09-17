package io

import (
	"context"
	"io"
	"strings"
)

type ctxInputReaderKey struct{}

// WithInputReader returns a copy of given context, with
// the input readcloser injected.
func WithInputReader(ctx context.Context, r io.ReadCloser) context.Context {
	parent := ctx
	if parent == nil {
		parent = context.Background()
	}

	return context.WithValue(parent, ctxInputReaderKey{}, r)
}

// InputReaderFromContext tries to extract the input reader from
// the given context. If not available, an empty reader would be returned.
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
