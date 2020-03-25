package db

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	tests := []struct {
		line     string
		expected string
		err      error
	}{
		{"EnD", "END", nil},
		{"?", "HELP", nil},
		{"blah", "blah", errors.New("Not a command. Type '?' for list of commands.")},
	}
	for _, test := range tests {
		output, err := Eval(test.line)

		assert.Equal(t, test.expected, output)
		assert.Equal(t, test.err, err)
	}
}
