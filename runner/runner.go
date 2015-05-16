package runner

import (
	"os/exec"
)

type Runner interface {
	Run(string, ...string) error
}

type run struct{}

func NewRunner() Runner {
	return &run{}
}

func (_ *run) Run(name string, args ...string) error {
	return exec.Command(name, args...).Run()
}
