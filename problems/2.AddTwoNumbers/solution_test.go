package problems

import (
	. "leetcode-go/adt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	l1 := new(LinkedList)
	l2 := new(LinkedList)
	l3 := new(LinkedList)

	l1.CreateList([]int{2, 4, 3})
	l2.CreateList([]int{5, 6, 4})
	l3.CreateList([]int{7, 0, 8})

	expected := l3.Head
	actual := addTwoNumbers(l1.Head, l2.Head)

	assert.Equal(t, expected, actual)
}
