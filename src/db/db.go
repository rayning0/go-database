package db

import (
	"errors"
	"fmt"
	"strings"
)

func get(words []string, m map[string]string) (string, error) {
	if len(words) != 2 {
		return "", errors.New("Invalid GET command. Format: GET [name]")
	}

	name := words[1]
	value, ok := m[name]

	if !ok {
		return "NULL", nil
	} else {
		return value, nil
	}
}

func set(words []string, m map[string]string) (string, error) {
	if len(words) != 3 {
		return "", errors.New("Invalid SET command. Format: SET [name] [value]")
	}
	name, value := words[1], words[2]

	m[name] = value
	value, ok := m[name]
	fmt.Println("m: ", m)

	if !ok {
		return "", errors.New("Error in setting " + name)
	} else {
		return value, nil
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
	case "set":
		return set(words, m)
	default:
		return line, errors.New("Invalid command. Type '?' for list of commands.")
	}
}
