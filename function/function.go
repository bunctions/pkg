package function

import "context"

// Callable defines the interface of a callable function.
type Callable interface {
	// Call is the method being invoke when the callable function being called.
	Call(context.Context) error
}

// SingleCallablePackager defines the interface of a packager that can
// generate a single callable function.
type SingleCallablePackager interface {
	// Pack generates a single callable function.
	Pack(args ...string) Callable
}

// MultiCallablePackager defines the interface of a packager that can
// generate multiple callable functions.
type MultiCallablePackager interface {
	// Pack generates a slice of callable functions.
	Pack(args ...string) []Callable
}
