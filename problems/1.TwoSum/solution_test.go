package problems

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	expected := []int{0, 1}
	actual := twoSum(nums, target)

	sort.Ints(actual)

	assert.Equal(t, expected, actual)
}
