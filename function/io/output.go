package io

import (
	"context"
	"io"
	"io/ioutil"
)

type ctxOutputWriterKey struct{}

// WithOutputWriter returns a copy of given context, with
// the output writer injected.
func WithOutputWriter(ctx context.Context, w io.Writer) context.Context {
	parent := ctx
	if parent == nil {
		parent = context.Background()
	}

	return context.WithValue(parent, ctxOutputWriterKey{}, w)
}

// OutputWriterFromContext tries to extract the output writer from
// the given context. If not available, an null writer would be returned.
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
