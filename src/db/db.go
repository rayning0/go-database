package db

import (
	"errors"
	"strings"
)

func get(words []string, m map[string]string) (string, error) {
	if len(words) != 2 {
		return "", errors.New("Invalid GET command. Format: GET [name]")
	}

	name := words[1]

	if m[name] == "" {
		return "NULL", nil
	} else {
		return m[name], nil
	}
}

func Eval(line string, m map[string]string) (string, error) {
	words := strings.Split(line, " ")
	word1 := strings.ToLower(words[0])

	switch word1 {
	case "end":
		return "END", nil
	case "?":
		return "HELP", nil
	case "get":
		return get(words, m)
	default:
		return line, errors.New("Invalid command. Type '?' for list of commands.")
	}
}
