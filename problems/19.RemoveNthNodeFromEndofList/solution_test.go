package problems

import (
	. "leetcode-go/adt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCases struct {
	head     *ListNode
	n        int
	expected *ListNode
}

func TestRemoveNthFromEnd(t *testing.T) {
	testCases := []TestCases{
		{
			head:     NewLinkedList([]int{1, 2, 3, 4, 5}),
			n:        2,
			expected: NewLinkedList([]int{1, 2, 3, 5}),
		},
		{
			head:     NewLinkedList([]int{1, 2}),
			n:        1,
			expected: NewLinkedList([]int{1}),
		},
		{
			head:     NewLinkedList([]int{1, 2, 3, 4}),
			n:        2,
			expected: NewLinkedList([]int{1, 2, 4}),
		},
		{
			head:     NewLinkedList([]int{1}),
			n:        1,
			expected: NewLinkedList([]int{}),
		},
		{
			head:     NewLinkedList([]int{1, 2, 3, 4, 5}),
			n:        1,
			expected: NewLinkedList([]int{1, 2, 3, 4}),
		},
	}

	for _, tc := range testCases {
		result := removeNthFromEnd(tc.head, tc.n)
		assert.Equal(t, tc.expected, result)
	}
}
