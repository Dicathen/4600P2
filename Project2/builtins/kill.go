package builtins

import (
	"errors"
	"strconv"
	"syscall"
)

func Kill(args ...string) error {
	if len(args) != 1 {
		return errors.New("kill: expected one argument")
	}
	pid, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}
	err = syscall.Kill(pid, syscall.SIGTERM)
	if err != nil {
		return err
	}
	return nil
}
