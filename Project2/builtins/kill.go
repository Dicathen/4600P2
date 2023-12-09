package builtins

import (
	"errors"
	"strconv"
	"syscall"
)

func Kill(args ...string) error {
	if len(args) > 1 {
		return errors.New("kill: expected one argument")
	}
	if len(args) == 0 {
		return errors.New("expected: kill <pid>")
	}
	if pid, err := strconv.Atoi(args[0]); err == nil {
		return syscall.Kill(pid, syscall.SIGTERM)
	}
	return nil
}
