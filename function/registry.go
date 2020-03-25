package function

import "errors"

var ErrNotFound = errors.New("not found")

type Registry interface {
	Register(Callable)
	Get() (Callable, bool)
	GetByName(string) (Callable, bool)
}

var DefaultRegistry Registry

func init() {
	DefaultRegistry = NewMemoryRegistry()
}

func Register(c Callable) {
	DefaultRegistry.Register(c)
}
