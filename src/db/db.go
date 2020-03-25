package db

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Main database map, name to value. Ex: {"a": "foo", "b": "foo"}
type M1 map[string]string

// Reverse of m1, mapping value to names. For SET and DELETE.
// Ex: {"foo": [a, b]}
type M2 map[string][]string

func get(words []string, m1 M1) (string, error) {
	if len(words) != 2 {
		return "", errors.New("Invalid GET command. Format: GET [name]")
	}
	name := words[1]

	if value, ok := m1[name]; ok {
		return value, nil
	} else {
		return "NULL", nil
	}
}

func set(words []string, m1 M1, m2 M2) error {
	if len(words) != 3 {
		return errors.New("Invalid SET command. Format: SET [name] [value]")
	}
	name, value := words[1], words[2]
	oldValue := m1[name]

	m1[name] = value
	_, ok1 := m1[name]
	fmt.Println("m1: ", m1)

	if _, ok2 := m2[oldValue]; ok2 {
		fmt.Println("m2["+oldValue+"] was: ", m2[oldValue])
		m2[oldValue] = deleteM1NameFromM2(name, m2[oldValue])
		fmt.Println("m2["+oldValue+"] now: ", m2[oldValue])
	}

	m2[value] = append(m2[value], name)
	value2 := m2[value][len(m2[value])-1]
	fmt.Println("m2: ", m2)

	if !ok1 || value2 != name {
		return errors.New("Error in setting " + name)
	} else {
		return nil
	}
}

func del(words []string, m1 M1, m2 M2) error {
	if len(words) != 2 {
		return errors.New("Invalid DELETE command. Format: DELETE [name]")
	}
	name := words[1]

	if value, ok := m1[name]; ok {
		delete(m1, name)
		fmt.Println("new m1: ", m1)
		m2[value] = deleteM1NameFromM2(name, m2[value])
		fmt.Println("new m2: ", m2)
		return nil
	} else {
		return errors.New("Can't delete " + name)
	}
}

// Deletes m1 name (example: a) from m2 (example: map[foo:[a b c]]),
// so it becomes map[foo:[b c]]
func deleteM1NameFromM2(name string, nameSlice []string) []string {
	for i, m1Name := range nameSlice {
		if m1Name == name {
			nameSlice = append(nameSlice[:i], nameSlice[i+1:]...)
			break
		}
	}
	return nameSlice
}

func count(words []string, m2 M2) (string, error) {
	if len(words) != 2 {
		return "", errors.New("Invalid COUNT command. Format: COUNT [value]")
	}
	value := words[1]

	return strconv.Itoa(len(m2[value])), nil
}

func begin(words []string, m1 M1, m2 M2, trans *Stack) error {
	if len(words) != 1 {
		return errors.New("Invalid BEGIN command. Format: BEGIN")
	}

	trans.Push(m1, m2)
	return nil
}

func rollback(words []string, m1 M1, m2 M2, trans *Stack) error {
	if len(words) != 1 {
		return errors.New("Invalid ROLLBACK command. Format: ROLLBACK")
	}

	element, ok := trans.Pop()
	fmt.Println("in rollback(): ", trans)
	if ok {
		fmt.Printf("%+v %+v\n", element.MainMap, element.ReverseMap)
		return nil
	} else {
		return errors.New("TRANSACTION NOT FOUND")
	}
}

// After COMMIT, may not ROLLBACK. Deletes all past transactions and makes fresh one.
func commit(words []string, m1 M1, m2 M2, trans *Stack) error {
	if len(words) != 1 {
		return errors.New("Invalid COMMIT command. Format: COMMIT")
	}
	var tnew *Stack

	trans = tnew
	trans.Push(m1, m2)
	return nil
}

func Eval(line string, m1 M1, m2 M2, trans Stack) (string, error) {
	words := strings.Split(line, " ")
	command := strings.ToLower(words[0])

	switch command {
	case "end":
		return "END", nil
	case "?":
		return "HELP", nil
	case "get":
		return get(words, m1)
	case "set":
		return "", set(words, m1, m2)
	case "count":
		return count(words, m2)
	case "delete":
		return "", del(words, m1, m2)
	case "begin":
		err := begin(words, m1, m2, &trans)
		fmt.Println("trans begin: ", trans)
		return "", err
	case "rollback":
		fmt.Println("trans before rollback: ", trans)
		err := rollback(words, m1, m2, &trans)
		fmt.Println("trans rollback: ", trans)
		return "", err
	case "commit":
		err := commit(words, m1, m2, &trans)
		fmt.Println("trans commit: ", trans)
		return "", err
	default:
		return line, errors.New("Invalid command. Type '?' for list of commands.")
	}
}
