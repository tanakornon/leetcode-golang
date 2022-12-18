package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCases struct {
	nums     []int
	expected [][]int
}

func TestThreeSum(t *testing.T) {
	testCases := []TestCases{
		{
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			nums:     []int{0, 0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			nums:     []int{1, 2, 3},
			expected: [][]int{},
		},
		{
			nums:     []int{-2, 0, 2},
			expected: [][]int{{-2, 0, 2}},
		},
		{
			nums:     []int{-2, -1, 0, 1, 2},
			expected: [][]int{{-2, 0, 2}, {-1, 0, 1}},
		},
	}

	for _, tc := range testCases {
		result := threeSum(tc.nums)
		assert.Equal(t, tc.expected, result)
	}
}
