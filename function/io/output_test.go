package io

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextWithOutputWriter(t *testing.T) {
	type testingFunc func(*testing.T)
	type args struct {
		ctx    context.Context
		writer io.Writer
	}
	type testData struct {
		data     args
		expected io.Writer
	}

	checking := func(d testData) testingFunc {
		return func(t *testing.T) {
			ctx := ContextWithOutputWriter(d.data.ctx, d.data.writer)

			assert.Equal(t, d.expected, ctx.Value(ctxOutputWriterKey{}))
		}
	}

	theWriter := bytes.NewBufferString("")

	testTable := map[string]testData{
		"BasicCase": {
			data: args{
				ctx:    context.Background(),
				writer: theWriter,
			},
			expected: theWriter,
		},
		"NilContextCase": {
			data: args{
				ctx:    nil,
				writer: theWriter,
			},
			expected: theWriter,
		},
	}

	for name, td := range testTable {
		t.Run(name, checking(td))
	}
}

func TestOutputWriterFromContext(t *testing.T) {
	type testingFunc func(*testing.T)
	type testData struct {
		data     context.Context
		expected io.Writer
	}

	checking := func(d testData) testingFunc {
		return func(t *testing.T) {
			w := OutputWriterFromContext(d.data)
			assert.Equal(t, d.expected, w)
		}
	}

	theWriter := bytes.NewBufferString("")

	testTable := map[string]testData{
		"BasicCase": {
			data: context.WithValue(
				context.Background(),
				ctxOutputWriterKey{},
				theWriter,
			),
			expected: theWriter,
		},
		"NilContextCase": {
			data:     nil,
			expected: ioutil.Discard,
		},
		"WrongTypeCase": {
			data: context.WithValue(
				context.Background(),
				ctxOutputWriterKey{},
				"Something Else",
			),
			expected: ioutil.Discard,
		},
	}

	for name, td := range testTable {
		t.Run(name, checking(td))
	}
}
