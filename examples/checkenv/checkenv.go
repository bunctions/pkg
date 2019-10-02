package main

import (
	"context"
	"encoding/json"

	"github.com/bunctions/pkg/function"
	funcio "github.com/bunctions/pkg/function/io"
)

var Exported = function.SingleCallablePackagerFunc(
	func(args ...string) function.Callable {
		return function.CallableFunc(checkenv)
	},
)

type response struct {
	Success     bool
	Environment map[string]interface{}
}

func checkenv(ctx context.Context) error {
	env := funcio.EnvironmentFromContext(ctx)
	writer := funcio.OutputWriterFromContext(ctx)

	resp := response{
		Success:     true,
		Environment: env,
	}

	encoder := json.NewEncoder(writer)
	return encoder.Encode(resp)
}
