package builtins

import (
	"errors"
	"fmt"
)

func Help(args ...string) error {
	switch len(args) {
	case 0:
		fmt.Println("cd: change directory")
		fmt.Println("echo: print arguments")
		fmt.Println("exit: exit shell")
		fmt.Println("help: print help")
		fmt.Println("pwd: print working directory")
		fmt.Println("kill: kill process")
		fmt.Println("history: print history")
		return nil
	default:
		return errors.New("help: expected zero arguments")
	}
}
