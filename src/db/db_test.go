package db

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	m1 := make(M1)
	m2 := make(M2)
	var trans Stack

	tests := []struct {
		line     string
		expected string
		err      error
	}{
		{"EnD", "END", nil},
		{"?", "HELP", nil},
		{"blah", "blah", errors.New("Invalid command. Type '?' for list of commands.")},
		{"GET a b", "", errors.New("Invalid GET command. Format: GET [name]")},
		{"GET a", "NULL", nil},
		{"SET a", "", errors.New("Invalid SET command. Format: SET [name] [value]")},
		{"SET a foo", "", nil},
		{"SET b foo", "", nil},
		{"SET c foo", "", nil},
		{"GET a", "foo", nil},
		{"GET b", "foo", nil},
		{"COUNT foo", "3", nil},
		{"COUNT bar", "0", nil},
		{"DELETE a", "", nil},
		{"GET a", "NULL", nil},
		{"COUNT foo", "2", nil},
		{"SET b baz", "", nil},
		{"COUNT foo", "1", nil},
		{"GET b", "baz", nil},
		{"GET B", "NULL", nil},
		{"eNd", "END", nil},
	}
	for _, test := range tests {
		output, err := Eval(test.line, &m1, &m2, &trans)

		assert.Equal(t, test.expected, output)
		assert.Equal(t, test.err, err)
	}
}

func TestExample2(t *testing.T) {
	m1 := make(M1)
	m2 := make(M2)
	var trans Stack

	tests := []struct {
		line     string
		expected string
		err      error
	}{
		{"SET a foo", "", nil},
		{"SET a foo", "", nil},
		{"COUNT foo", "1", nil},
		{"GET a", "foo", nil},
		{"DELETE a", "", nil},
		{"GET a", "NULL", nil},
		{"COUNT foo", "0", nil},
		{"END", "END", nil},
	}
	for _, test := range tests {
		output, err := Eval(test.line, &m1, &m2, &trans)

		assert.Equal(t, test.expected, output)
		assert.Equal(t, test.err, err)
	}
}

func TestExample3(t *testing.T) {
	m1 := make(M1)
	m2 := make(M2)
	var trans Stack

	tests := []struct {
		line     string
		expected string
		err      error
	}{
		{"ROLLBACK", "", errors.New("TRANSACTION NOT FOUND")},
		{"BEGIN", "", nil},
		{"SET a foo", "", nil},
		{"GET a", "foo", nil},
		{"BEGIN", "", nil},
		{"SET a bar", "", nil},
		{"GET a", "bar", nil},
		{"SET a baz", "", nil},
		{"ROLLBACK", "", nil},
		{"GET a", "foo", nil},
		{"ROLLBACK", "", nil},
		{"GET a", "NULL", nil},
		{"END", "END", nil},
	}
	for _, test := range tests {
		output, err := Eval(test.line, &m1, &m2, &trans)

		assert.Equal(t, test.expected, output)
		assert.Equal(t, test.err, err)
	}
}
