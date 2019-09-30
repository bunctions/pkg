package main

import "github.com/bunctions/pkg/runner"

type config struct {
	*runner.Config

	Port uint `default:"8080"`
}
