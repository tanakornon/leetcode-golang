package problems

import (
	. "leetcode-go/adt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	l1 := NewLinkedList([]int{2, 4, 3})
	l2 := NewLinkedList([]int{5, 6, 4})

	expected := NewLinkedList([]int{7, 0, 8})
	actual := addTwoNumbers(l1, l2)

	assert.Equal(t, expected, actual)
}
