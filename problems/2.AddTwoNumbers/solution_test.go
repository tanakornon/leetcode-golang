package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	l1 := new(LinkedList)
	l2 := new(LinkedList)
	l3 := new(LinkedList)

	l1.createList([]int{2, 4, 3})
	l2.createList([]int{5, 6, 4})
	l3.createList([]int{7, 0, 8})

	expected := l3.head
	actual := addTwoNumbers(l1.head, l2.head)

	assert.Equal(t, expected, actual)
}
