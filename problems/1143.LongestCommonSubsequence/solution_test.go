package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	text1 := "abcde"
	text2 := "ace"

	expected := 3
	actual := longestCommonSubsequence(text1, text2)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	text1 := "abc"
	text2 := "abc"

	expected := 3
	actual := longestCommonSubsequence(text1, text2)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	text1 := "abc"
	text2 := "def"

	expected := 0
	actual := longestCommonSubsequence(text1, text2)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	text1 := "bsbininm"
	text2 := "jmjkbkjkv"

	expected := 1
	actual := longestCommonSubsequence(text1, text2)

	assert.Equal(t, expected, actual)
}
