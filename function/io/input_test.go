package io

import (
	"context"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextWithInputReader(t *testing.T) {
	type testingFunc func(*testing.T)
	type args struct {
		ctx        context.Context
		readCloser io.ReadCloser
	}
	type testData struct {
		data     args
		expected io.Reader
	}

	checking := func(d testData) testingFunc {
		return func(t *testing.T) {
			ctx := ContextWithInputReader(d.data.ctx, d.data.readCloser)

			assert.Equal(
				t,
				d.expected,
				ctx.Value(ctxInputReaderKey{}),
			)
		}
	}

	theReadCloser := ioutil.NopCloser(strings.NewReader("Hello World"))

	testTable := map[string]testData{
		"BasicCase": {
			data: args{
				ctx:        context.Background(),
				readCloser: theReadCloser,
			},
			expected: theReadCloser,
		},
		"NilContextCase": {
			data: args{
				ctx:        nil,
				readCloser: theReadCloser,
			},
			expected: theReadCloser,
		},
	}

	for name, td := range testTable {
		t.Run(name, checking(td))
	}
}

func TestInputReaderFromContext(t *testing.T) {
	type testingFunc func(*testing.T)
	type expectation struct {
		success bool
		reader  io.Reader
	}
	type testData struct {
		data     context.Context
		expected expectation
	}

	checking := func(d testData) testingFunc {
		return func(t *testing.T) {
			r := InputReaderFromContext(d.data)

			if d.expected.success {
				assert.Equal(t, d.expected.reader, r)
				return
			}

			// checking whether it is emptyReader
			er, ok := r.(*strings.Reader)
			assert.True(t, ok)
			assert.Equal(
				t,
				*strings.NewReader(""),
				*er,
			)
		}
	}

	theReadCloser := ioutil.NopCloser(strings.NewReader("Hello World"))

	testTable := map[string]testData{
		"BasicCase": {
			data: context.WithValue(
				context.Background(),
				ctxInputReaderKey{},
				theReadCloser,
			),
			expected: expectation{
				success: true,
				reader:  theReadCloser,
			},
		},
		"NilContextCase": {
			data: nil,
			expected: expectation{
				success: false,
				reader:  nil,
			},
		},
		"WrongTypeCase": {
			data: context.WithValue(
				context.Background(),
				ctxInputReaderKey{},
				"Something Else",
			),
			expected: expectation{
				success: false,
				reader:  nil,
			},
		},
	}

	for name, td := range testTable {
		t.Run(name, checking(td))
	}
}
