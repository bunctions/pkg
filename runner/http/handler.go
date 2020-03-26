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
			ctx := funcio.ContextWithInputReader(r.Context(), r.Body)
			ctx = funcio.ContextWithOutputWriter(ctx, rw)
			ctx = funcio.ContextWithEnvironments(
				ctx,
				runnerutil.GetSystemEnvironment(),
				getRequestHeaderAsEnvironment(r),
				getRequestParamsAsEnvironment(r),
			)

			_ = callable.Call(ctx)
		},
	)
}
