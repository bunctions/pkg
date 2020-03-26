package http

import (
	"errors"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/bunctions/pkg/function"
	"github.com/bunctions/pkg/function/logger"
	"go.uber.org/zap"
)

type router http.Handler

/**
 * pathRouter serve 2 kind of paths:
 *   - serve `/` using defaultHandler
 *   - serve `/functions/<func_name>` using handlers with
 *     the look up using function name
 */
type pathRouter struct {
	defaultHandler http.Handler
	handlers       map[string]http.Handler
}

func newPathRouter(registry function.Registry, logger *zap.Logger) router {
	prtr := &pathRouter{
		defaultHandler: http.NotFoundHandler(),
		handlers:       map[string]http.Handler{},
	}

	if c, ok := registry.Get(); ok {
		logger.Info("setup default handler")
		prtr.defaultHandler = newHandler(c)
	}

	cs := registry.GetAll()
	for _, c := range cs {
		if nc, ok := c.(function.NamedCallable); ok {
			name := escapeFuncName(nc.GetName())
			logger.Info("setup handler", zap.String("name", name))
			prtr.handlers[name] = newHandler(nc)
		}
	}

	return prtr
}

func (prtr *pathRouter) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	logger := logger.LoggerFromContext(r.Context())

	var err error
	defer func() {
		if err != nil {
			logger.Warn("error handling", zap.Error(err))
			http.NotFound(rw, r)
			return
		}
	}()

	path := strings.ToLower(r.URL.Path)
	path = strings.TrimLeft(path, "/")
	path = strings.TrimRight(path, "/")
	if path == "" {
		prtr.defaultHandler.ServeHTTP(rw, r)
		return
	}

	pathSplits := strings.SplitN(
		strings.TrimRight(path, "/"),
		"/",
		2,
	)

	if len(pathSplits) != 2 {
		err = errors.New("path mismatch: split != 2")
		return
	}

	if pathSplits[0] != "functions" {
		err = errors.New("path mismatch: not prefixed /funcsions")
		return
	}

	target, err := url.PathUnescape(pathSplits[1])
	if err != nil {
		err = errors.New("path mismatch: fail to unescape")
		return
	}

	handler, ok := prtr.handlers[escapeFuncName(target)]
	if !ok {
		err = errors.New("path mismatch: handler not found")
		return
	}

	handler.ServeHTTP(rw, r)
}

func escapeFuncName(name string) string {
	rgx := regexp.MustCompile("[^a-z]")
	result := rgx.ReplaceAll([]byte(strings.ToLower(name)), []byte("_"))
	return string(result)
}
