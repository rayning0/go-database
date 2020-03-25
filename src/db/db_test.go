package db

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	m := make(map[string]string)

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
		{"SET a foo", "foo", nil},
		{"SET b foo", "foo", nil},
		{"GET a", "foo", nil},
		{"GET b", "fool", nil},
	}
	for _, test := range tests {
		output, err := Eval(test.line, m)

		assert.Equal(t, test.expected, output)
		assert.Equal(t, test.err, err)
	}
}
