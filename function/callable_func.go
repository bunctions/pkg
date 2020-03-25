package function

import "context"

// CallableFunc wraps a function as a Callable.
type CallableFunc func(context.Context) error

func (f CallableFunc) Call(ctx context.Context) error {
	return f(ctx)
}

type namedCallable struct {
	name     string
	callable Callable
}

func NewNamedCallable(name string, callable Callable) NamedCallable {
	return namedCallable{
		name:     name,
		callable: callable,
	}
}

func (nc namedCallable) GetName() string {
	return nc.name
}

func (nc namedCallable) Call(ctx context.Context) error {
	return nc.callable.Call(ctx)
}
