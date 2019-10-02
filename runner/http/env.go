package main

import (
	"net/http"
	"strings"

	funcio "github.com/bunctions/pkg/function/io"
)

func getRequestHeaderAsEnvironment(r *http.Request) funcio.Environment {
	env := funcio.Environment{}

	for key, values := range r.Header {
		// only respect the first occurrence
		env[strings.ToLower(key)] = values[0]
	}

	return env
}

func getRequestParamsAsEnvironment(r *http.Request) funcio.Environment {
	env := funcio.Environment{}

	for key, values := range r.URL.Query() {
		// only respect the first occurrence
		env[strings.ToLower(key)] = values[0]
	}

	return env
}
