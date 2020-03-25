package db

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func get(words []string, m1 map[string]string) (string, error) {
	if len(words) != 2 {
		return "", errors.New("Invalid GET command. Format: GET [name]")
	}
	name := words[1]
	value, ok := m1[name]

	if !ok {
		return "NULL", nil
	} else {
		return value, nil
	}
}

func set(words []string, m1 map[string]string, m2 map[string][]string) (string, error) {
	if len(words) != 3 {
		return "", errors.New("Invalid SET command. Format: SET [name] [value]")
	}
	name, value := words[1], words[2]

	m1[name] = value
	value1, ok1 := m1[name]
	fmt.Println("m1: ", m1)

	m2[value] = append(m2[value], name)
	value2 := m2[value][len(m2[value])-1]
	fmt.Println("m2: ", m2)

	if !ok1 || value2 != name {
		return "", errors.New("Error in setting " + name)
	} else {
		return value1, nil
	}
}

func count(words []string, m2 map[string][]string) (string, error) {
	if len(words) != 2 {
		return "", errors.New("Invalid COUNT command. Format: COUNT [value]")
	}
	value := words[1]

	return strconv.Itoa(len(m2[value])), nil
}

func Eval(line string, m1 map[string]string, m2 map[string][]string) (string, error) {
	words := strings.Split(line, " ")
	word1 := strings.ToLower(words[0])

	switch word1 {
	case "end":
		return "END", nil
	case "?":
		return "HELP", nil
	case "get":
		return get(words, m1)
	case "set":
		return set(words, m1, m2)
	case "count":
		return count(words, m2)
	default:
		return line, errors.New("Invalid command. Type '?' for list of commands.")
	}
}
