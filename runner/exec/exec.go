package exec

import (
	"go.uber.org/zap"

	"github.com/bunctions/pkg/runner/util"
)

const Name = "exec"

func Runner() execRunner {
	return execRunner{}
}

type execRunner struct{}

func (execRunner) Start() {
	_ = util.NewLogger().
		With(zap.String("runner", "http"))
}
