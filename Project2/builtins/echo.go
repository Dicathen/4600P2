package builtins

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func Echo(args ...string) error {
	if len(args) == 0 {
		return errors.New("echo: expected at least one argument")
	} else if len(args) == 1 && args[0] == "*" {
		//print files in current directory

		//get current directory
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		//get files in current directory
		files, err := os.ReadDir(wd)
		if err != nil {
			return err
		}

		//print files in current directory in a single line
		for _, file := range files {
			fmt.Print(file.Name() + " ")
		}
		fmt.Println()

		return nil
	}
	fmt.Println(strings.Join(args, " "))
	return nil
}
