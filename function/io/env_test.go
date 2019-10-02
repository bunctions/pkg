package io

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithEnvironment(t *testing.T) {
	type testingFunc func(*testing.T)
	type args struct {
		ctx context.Context
		env Environment
	}
	type testData struct {
		data     args
		expected Environment
	}

	checking := func(d testData) testingFunc {
		return func(t *testing.T) {
			ctx := WithEnvironment(d.data.ctx, d.data.env)

			assert.Equal(
				t,
				d.expected,
				ctx.Value(ctxEnvironmentKey{}),
			)
		}
	}

	theEnv := Environment{
		"hello":       "world",
		"foo":         "bar",
		"okay-google": 123,
	}

	testTable := map[string]testData{
		"BasicCase": {
			data: args{
				ctx: context.Background(),
				env: theEnv,
			},
			expected: theEnv,
		},
		"NilContextCase": {
			data: args{
				ctx: nil,
				env: theEnv,
			},
			expected: theEnv,
		},
	}

	for name, td := range testTable {
		t.Run(name, checking(td))
	}
}

func TestEnvironmentFromContext(t *testing.T) {
	type testingFunc func(*testing.T)
	type testData struct {
		data     context.Context
		expected Environment
	}

	checking := func(d testData) testingFunc {
		return func(t *testing.T) {
			assert.Equal(t, d.expected, EnvironmentFromContext(d.data))
		}
	}

	theEnv := Environment{
		"hello":       "world",
		"foo":         "bar",
		"okay-google": 123,
	}

	testTable := map[string]testData{
		"BasicCase": {
			data: context.WithValue(
				context.Background(),
				ctxEnvironmentKey{},
				theEnv,
			),
			expected: theEnv,
		},
		"NilContextCase": {
			data:     nil,
			expected: Environment{},
		},
		"WrongTypeCase": {
			data: context.WithValue(
				context.Background(),
				ctxEnvironmentKey{},
				"Something Else",
			),
			expected: Environment{},
		},
	}

	for name, td := range testTable {
		t.Run(name, checking(td))
	}
}
