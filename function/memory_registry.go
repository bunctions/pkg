package function

const (
	unnamedCallableKey = "__unnamed__"
)

type memoryRegistry map[string]Callable

func NewMemoryRegistry() Registry {
	return memoryRegistry{}
}

func (mr memoryRegistry) Register(c Callable) {
	key := unnamedCallableKey
	if nc, ok := c.(NamedCallable); ok {
		key = nc.GetName()
	}

	mr[key] = c
}

func (mr memoryRegistry) Get() (Callable, bool) {
	if len(mr) == 1 {
		// if only 1 callable got registered, return it
		for _, c := range mr {
			return c, true
		}
	}

	return mr.GetByName(unnamedCallableKey)

}

func (mr memoryRegistry) GetByName(string) (Callable, bool) {
	c, ok := mr[unnamedCallableKey]
	return c, ok
}
