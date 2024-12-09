package devenv

import (
	"os/exec"
	"syscall"

	"github.com/juju/errors"
)

func IsExitStatus(err error, status int) bool {
	var exitErr *exec.ExitError
	if errors.As(errors.Cause(err), &exitErr) {
		return exitErr.Sys().(syscall.WaitStatus).ExitStatus() == status
	}

	return false
}
