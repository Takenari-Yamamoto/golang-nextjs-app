package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusSimply(t *testing.T) {
	result := Plus(1, 3)
	expect := 4

	assert.Equal(t, expect, result)
}

func TestPlusAssertFailContinue(t *testing.T) {
	result1 := Plus(1, 3)
	expect1 := 4

	assert.Equal(t, expect1, result1)

	result2 := Plus(1, 3)
	badExpect2 := 4 // miss

	assert.Equal(t, badExpect2, result2)

	result3 := Plus(3, 5)
	expect3 := 8

	assert.Equal(t, expect3, result3, "not equal")

	result4 := Plus(3, 5)
	badExpect4 := 8 // miss

	assert.Equal(t, badExpect4, result4, "not equal")
}
