package util

import (
	"os"
	"strings"

	funcio "github.com/bunctions/pkg/function/io"
)

// GetSystemEnvironment returns the system environment
func GetSystemEnvironment() funcio.Environment {
	env := funcio.Environment{}
	envList := os.Environ()

	for _, eachEnv := range envList {
		splits := strings.Split(eachEnv, "=")
		key := strings.ToLower(splits[0])

		if len(splits) < 2 {
			env[key] = ""
			continue
		}

		env[key] = strings.Join(splits[1:], "=")
	}

	return env
}
