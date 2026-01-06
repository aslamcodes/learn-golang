package main

import (
	"bytes"
	"os/exec"
)

type execution_step struct {
	step
}

func (s execution_step) execute() (string, error) {
	out := bytes.Buffer{}

	cmd := exec.Command(s.exe, s.args...)

	cmd.Stdout = &out
	cmd.Dir = s.project

	if err := cmd.Run(); err != nil {
		return "", &stepErr{
			step:  s.name,
			msg:   "failed to execute",
			cause: err,
		}
	}

	if out.Len() > 0 {
		return "", &stepErr{
			step:  s.name,
			msg:   "unexpected output",
			cause: nil,
		}
	}

	return s.message, nil
}

func NewExecutionStep(name, project, message, exe string, args []string) execution_step {
	s := execution_step{}

	s.step = newStep(name, project, message, exe, args)

	return s
}
