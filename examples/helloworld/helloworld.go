package main

import (
	"context"
	"io"

	"github.com/bunctions/pkg/function"
	funcio "github.com/bunctions/pkg/function/io"
)

var Exported = function.SingleCallablePackagerFunc(
	func(args ...string) function.Callable {
		return function.CallableFunc(helloworld)
	},
)

func helloworld(ctx context.Context) error {
	writer := funcio.OutputWriterFromContext(ctx)
	_, err := io.WriteString(writer, "hello world")

	return err
}
