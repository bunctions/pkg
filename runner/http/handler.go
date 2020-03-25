package http

import (
	"net/http"

	"github.com/bunctions/pkg/function"
	funcio "github.com/bunctions/pkg/function/io"
	runnerutil "github.com/bunctions/pkg/runner/util"
)

func newHandler(callable function.Callable) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			ctx := funcio.WithInputReader(r.Context(), r.Body)
			ctx = funcio.WithOutputWriter(ctx, rw)
			ctx = funcio.WithEnvironments(
				ctx,
				runnerutil.GetSystemEnvironment(),
				getRequestHeaderAsEnvironment(r),
				getRequestParamsAsEnvironment(r),
			)

			_ = callable.Call(ctx)
		},
	)
}
