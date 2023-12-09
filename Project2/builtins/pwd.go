package builtins

import (
	"errors"
	"fmt"
	"os"
)


func PrintWorkingDirectory(args ...string) error {
	switch len(args) {
	case 0:
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		fmt.Println(wd)
		return nil
	default:
		return errors.New("pwd: expected zero arguments")
	}
}