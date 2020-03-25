package db

import (
	"errors"
	"strings"
)

func Eval(line string) (string, error) {
	if strings.ToLower(line) == "end" {
		return "END", nil
	} else if line == "?" {
		return "HELP", nil
	} else {
		return line, errors.New("Not a command. Type '?' for list of commands.")
	}
}
