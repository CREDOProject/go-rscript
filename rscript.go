package gorscript

import (
	"fmt"
	"os/exec"
)

type rscript struct {
	binaryName string
	args       []string
}

// New retuns a new rscript object.
func New(binaryName string) *rscript {
	o := &rscript{
		binaryName: binaryName,
		args:       []string{},
	}
	return o
}

// Evaluate adds a new command to evaluate.
func (o *rscript) Evaluate(script string) *rscript {
	o.args = append(o.args, "-e", script)
	return o
}

// Seal command to be run.
func (o *rscript) Seal() (*exec.Cmd, error) {
	if o.binaryName == "" {
		return nil, fmt.Errorf("binaryName not provided.")
	}
	if len(o.args) < 1 {
		return nil, fmt.Errorf("Scripts to evaluate were not provided.")
	}
	return execCommander().Command(o.binaryName, o.args...), nil
}
