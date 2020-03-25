package runner

import (
	"log"
	"strings"

	runnerhttp "github.com/bunctions/pkg/runner/http"
	"github.com/kelseyhightower/envconfig"
)

// Start starts a runner base on the config
func Start() {
	conf := &config{}
	err := envconfig.Process("", conf)
	if err != nil {
		log.Fatalf("Error: %s", err)
		return
	}

	runnerType := strings.ToLower(conf.Runner)
	switch runnerType {
	case "http":
		runnerhttp.Start()
	default:
		log.Fatalf("Error: unknown runner \"%s\"", conf.Runner)
		return
	}
}
