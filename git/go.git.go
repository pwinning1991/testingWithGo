package git

import (
	"os/exec"
	"strings"
)

var execCommand = exec.Command

func Version() string {
	cmd := execCommand("git", "version")
	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	n := len("git version ")
	version := string(stdout[n:])
	return strings.TrimSpace(version)
}

type Checker struct {
	execCommand func(name string, args ...string) *exec.Cmd
}

func (gc *Checker) command(name string, args ...string) *exec.Cmd {
	if gc.execCommand == nil {
		return execCommand(name, args...)
	}
	return gc.execCommand(name, args...)
}

func (gc *Checker) Version() string {
	cmd := gc.execCommand("git", "version")
	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	n := len("git version ")
	version := string(stdout[n:])
	return strings.TrimSpace(version)
}
