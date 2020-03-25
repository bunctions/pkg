package function

import "context"

// Callable defines the interface of a callable function.
type Callable interface {
	// Call is the method being invoke when the callable function being called.
	Call(context.Context) error
}

// NamedCallable defines the interface of a named callable function.
type NamedCallable interface {
	Callable

	// GetName returns the name of the callable function.
	GetName() string
}
