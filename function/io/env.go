package io

import (
	"context"
)

type ctxEnvironmentKey struct{}

// Environment defines the representation of the execution environment.
type Environment map[string]interface{}

// ContextWithEnvironment returns a copy of given context, with
// the environment injected.
func ContextWithEnvironment(ctx context.Context, env Environment) context.Context {
	return ContextWithEnvironments(ctx, env)
}

// ContextWithEnvironments returns a copy of given context, with
// the multiple environments injected.
func ContextWithEnvironments(ctx context.Context, envs ...Environment) context.Context {
	parent := ctx
	if parent == nil {
		parent = context.Background()
	}

	env := MergeEnvironments(envs...)

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

func MergeEnvironments(envs ...Environment) Environment {
	env := Environment{}

	for _, eachEnv := range envs {
		for key, value := range eachEnv {
			env[key] = value
		}
	}

	return env
}
