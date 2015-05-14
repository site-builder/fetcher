package commander

import (
	"os/exec"
)

type Command interface {
	Run() error
}

type command struct{}

func CreateCommand(name string, arg ...string) Command {
	return exec.Command(name, arg...)
}

func (command *command) Run() error {
	return command.Run()
}
