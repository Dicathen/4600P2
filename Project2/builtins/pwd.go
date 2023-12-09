package builtins

import (
	"errors"
	"fmt"
	"os"
)

func PrintWorkingDirectory(args ...string) error {
	switch len(args) {
	case 0:
		fmt.Println(os.Getwd())
		return nil
	default:
		return errors.New("pwd: expected zero arguments")
	}
}
