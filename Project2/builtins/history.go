package builtins

import (
	"errors"
	"fmt"
)

//create a slice of strings to hold the history
var historyList []string

func History(args ...string) error {
	switch len(args) {
	case 0:
		for i, cmd := range historyList {
			fmt.Printf("%d\t%s\n", i+1, cmd)
		}
		return nil
	default:
		return errors.New("history: expected zero arguments")
	}
}

//append to historyList
func SetHistoryList(input string) {
	historyList = append(historyList, input)
}
