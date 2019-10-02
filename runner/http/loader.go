package main

import (
	"errors"
	"net/http"
	"plugin"

	"github.com/bunctions/pkg/function"
	funcio "github.com/bunctions/pkg/function/io"
)

type loader struct {
	path   string
	symbol string
}

var ErrUnknownSymbolType = errors.New("Unknown symbol type")

func (l *loader) loadHandler() (http.Handler, error) {
	loaded, err := plugin.Open(l.path)
	if err != nil {
		return nil, err
	}

	symbol, err := loaded.Lookup(l.symbol)
	if err != nil {
		return nil, err
	}

	handler := http.Handler(nil)

	switch s := symbol.(type) {
	case function.SingleCallablePackager:
		handler = l.handleSingleCallable(s)
	case function.MultiCallablePackager:
		handler = l.handleMultipleCallable(s)
	default:
		return nil, ErrUnknownSymbolType
	}

	return handler, nil
}

func (l *loader) handleSingleCallable(
	p function.SingleCallablePackager,
) http.HandlerFunc {
	callable := p.Pack()

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ctx := funcio.WithInputReader(r.Context(), r.Body)
		ctx = funcio.WithOutputWriter(ctx, rw)

		_ = callable.Call(ctx)
	})
}

func (l *loader) handleMultipleCallable(
	m function.MultiCallablePackager,
) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusNotImplemented)
	})
}
