package main

import (
	"context"
	"encoding/json"

	"github.com/bunctions/pkg/function"
	funcio "github.com/bunctions/pkg/function/io"
	"github.com/bunctions/pkg/runner"
)

func init() {
	function.RegisterNamedFunc("check_env", checkenv)
}

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

func main() {
	runner.Start()
}
