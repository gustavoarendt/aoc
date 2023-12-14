package day01

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSteps(t *testing.T) {
	tests := []struct {
		expected int
		input    string
		fn     func(string) int
	}{
		{
			expected: 142,
			input:    "test.txt",
			fn: stepOne,
		},
		{
			expected: 281,
			input:    "test2.txt",
			fn: stepTwo,
		},
	}

	for _, test := range tests {
		content, err := os.ReadFile(test.input)
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, test.fn(string(content)))
	}
}