package runner

import (
	"strings"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"

	"github.com/bunctions/pkg/runner/exec"
	"github.com/bunctions/pkg/runner/http"
	"github.com/bunctions/pkg/runner/util"
)

// Generic is a generic runner which will choose a proper specific runner
// to start.
var Generic genericRunner

type genericRunner struct {
}

type config struct {
	Runner string `split_words:"true" default:"http"`
}

func (genericRunner) Start() {
	logger := util.NewLogger().
		With(zap.String("runner", "generic"))

	conf := &config{}
	err := envconfig.Process("", conf)
	if err != nil {
		logger.Error(
			"failed to process config",
			zap.Error(err),
		)
		return
	}

	runnerType := strings.ToLower(conf.Runner)
	logger = logger.With(zap.String(
		"desired_runner",
		runnerType,
	))

	r := Runner(nil)
	switch runnerType {
	case http.Name:
		r = http.Runner()
	case exec.Name:
		r = exec.Runner()
	default:
		logger.Error("failed to detect the runner")
		return
	}

	r.Start()
}
