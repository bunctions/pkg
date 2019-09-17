package io

import (
	"context"
	"io"
	"io/ioutil"
)

type ctxOutputWriterKey struct{}

func WithOutputWriter(ctx context.Context, w io.Writer) context.Context {
	parent := ctx
	if parent == nil {
		parent = context.Background()
	}

	return context.WithValue(parent, ctxOutputWriterKey{}, w)
}

func OutputWriterFromContext(ctx context.Context) io.Writer {
	nullWriter := ioutil.Discard
	if ctx == nil {
		return nullWriter
	}

	w, ok := ctx.Value(ctxOutputWriterKey{}).(io.Writer)
	if !ok {
		return nullWriter
	}

	return w
}
