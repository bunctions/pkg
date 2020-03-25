package main

import (
	"context"
	"io"

	"github.com/bunctions/pkg/function"
	funcio "github.com/bunctions/pkg/function/io"
	"github.com/bunctions/pkg/runner"
)

func init() {
	function.Register(function.CallableFunc(helloworld))
}

func helloworld(ctx context.Context) error {
	writer := funcio.OutputWriterFromContext(ctx)
	_, err := io.WriteString(writer, "hello world")

	return err
}

func main() {
	runner.Start()
}
