package main

import (
	"os/exec"
)

type step struct {
	name    string
	exe     string
	args    []string
	message string
	project string
}

func (s step) execute() (string, error) {
	cmd := exec.Command(s.exe, s.args...)

	cmd.Dir = s.project

	if err := cmd.Run(); err != nil {
		return "", &stepErr{
			step:  s.name,
			msg:   "failed to execute",
			cause: err,
		}
	}

	return s.message, nil
}

func newStep(name, project, message, exe string, args []string) step {
	return step{
		name:    name,
		project: project,
		message: message,
		exe:     exe,
		args:    args,
	}
}
