package function

import "errors"

var ErrNotFound = errors.New("not found")

type Registry interface {
	Register(Callable)
	Get() (Callable, bool)
	GetByName(string) (Callable, bool)
	GetAll() []Callable
}

var DefaultRegistry Registry

func init() {
	DefaultRegistry = NewMemoryRegistry()
}

// Register register a callable to default register.
func Register(c Callable) {
	DefaultRegistry.Register(c)
}

// RegisterFunc register a callable function to
// default register.
func RegisterFunc(f CallableFunc) {
	DefaultRegistry.Register(f)
}

// RegisterNamed register a callable with a name to
// default register.
func RegisterNamed(name string, c Callable) {
	Register(NewNamedCallable(name, c))
}

// RegisterNamedFunc register a callable function
// with a name to default register.
func RegisterNamedFunc(name string, f CallableFunc) {
	Register(NewNamedCallable(name, f))
}
