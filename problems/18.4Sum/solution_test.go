package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCases struct {
	nums     []int
	target   int
	expected [][]int
}

func TestFourSum(t *testing.T) {
	testCases := []TestCases{
		{
			nums:     []int{1, 0, -1, 0, -2, 2},
			target:   0,
			expected: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			nums:     []int{-3, -2, -1, 0, 0, 1, 2, 3},
			target:   0,
			expected: [][]int{{-3, -2, 2, 3}, {-3, -1, 1, 3}, {-3, 0, 0, 3}, {-3, 0, 1, 2}, {-2, -1, 0, 3}, {-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			nums:     []int{},
			target:   0,
			expected: [][]int{},
		},
		{
			nums:     []int{0, 0, 0, 0},
			target:   0,
			expected: [][]int{{0, 0, 0, 0}},
		},
	}

	for _, tc := range testCases {
		result := fourSum(tc.nums, tc.target)
		assert.Equal(t, tc.expected, result)
	}
}
