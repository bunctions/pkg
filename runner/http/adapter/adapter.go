package adapter

import "net/http"

type Adapter interface {
	Apply(http.Handler) http.Handler
}

func ApplyAll(h http.Handler, adapters ...Adapter) http.Handler {
	finalHandler := h
	for _, eachAdapter := range adapters {
		finalHandler = eachAdapter.Apply(finalHandler)
	}

	return finalHandler
}

type AdapterFunc func(http.Handler) http.Handler

func (f AdapterFunc) Apply(next http.Handler) http.Handler {
	return f(next)
}
