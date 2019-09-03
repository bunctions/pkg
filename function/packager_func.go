package function

// SingleCallablePackagerFunc wraps a function as
// a SingleCallablePackager.
type SingleCallablePackagerFunc func(...string) Callable

func (p SingleCallablePackagerFunc) Pack(args ...string) Callable {
	return p(args...)
}

// MultiCallablePackagerFunc wraps a function as
// a MultiCallablePackager.
type MultiCallablePackagerFunc func(...string) []Callable

func (p MultiCallablePackagerFunc) Pack(args ...string) []Callable {
	return p(args...)
}
