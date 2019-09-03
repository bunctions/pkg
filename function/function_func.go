package function

import "context"

// CallableFunc wraps a function as a Callable.
type CallableFunc func(context.Context) error

func (f CallableFunc) Call(ctx context.Context) error {
	return f(ctx)
}
