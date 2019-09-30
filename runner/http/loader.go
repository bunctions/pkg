package main

import (
	"errors"
	"net/http"
	"plugin"

	"github.com/bunctions/pkg/function"
	funcio "github.com/bunctions/pkg/function/io"
)

type loader struct {
	path    string
	symbol  string
	handler http.HandlerFunc
}

var ErrUnknownSymbolType = errors.New("Unknown symbol type")

func (l *loader) load() error {
	loaded, err := plugin.Open(l.path)
	if err != nil {
		return err
	}

	symbol, err := loaded.Lookup(l.symbol)
	if err != nil {
		return err
	}

	switch s := symbol.(type) {
	case function.SingleCallablePackager:
		l.handler = l.handleSingleCallable(s)
	case function.MultiCallablePackager:
		l.handler = l.handleMultipleCallable(s)
	default:
		return ErrUnknownSymbolType
	}

	return nil
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
