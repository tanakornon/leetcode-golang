package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCases struct {
	tokens   []string
	expected int
}

func TestEvalRPN(t *testing.T) {
	testCases := []TestCases{
		{
			tokens:   []string{"2", "1", "+", "3", "*"},
			expected: 9,
		},
		{
			tokens:   []string{"4", "13", "5", "/", "+"},
			expected: 6,
		},
		{
			tokens:   []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			expected: 22,
		},
		{
			tokens:   []string{"2", "1", "3", "*", "+"},
			expected: 5,
		},
		{
			tokens:   []string{"5", "1", "2", "+", "4", "*", "+", "3", "-"},
			expected: 14,
		},
	}

	for _, tc := range testCases {
		result := evalRPN(tc.tokens)
		assert.Equal(t, tc.expected, result)
	}
}
