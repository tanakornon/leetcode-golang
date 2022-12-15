package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	s := []byte("hello")

	reverseString(s)

	expected := "olleh"
	actual := string(s)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	s := []byte("Hannah")

	reverseString(s)

	expected := "hannaH"
	actual := string(s)

	assert.Equal(t, expected, actual)
}
