package io

import (
	"context"
)

type ctxEnvironmentKey struct{}

// Environment defines the representation of the execution environment.
type Environment map[string]interface{}

// WithEnvironment returns a copy of given context, with
// the environment injected.
func WithEnvironment(ctx context.Context, env Environment) context.Context {
	parent := ctx
	if parent == nil {
		parent = context.Background()
	}

	return context.WithValue(parent, ctxEnvironmentKey{}, env)
}

// EnvironmentFromContext tries to extract the environment reader from
// the given context. If not available, an empty environment would be returned.
func EnvironmentFromContext(ctx context.Context) Environment {
	emptyEnv := Environment{}
	if ctx == nil {
		return emptyEnv
	}

	env, ok := ctx.Value(ctxEnvironmentKey{}).(Environment)
	if !ok {
		return emptyEnv
	}

	return env
}
