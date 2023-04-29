package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert := assert.New(t)

	result := Add(1, 2)
	expected := 3
	assert.Equal(expected, result, "1 + 2 は 3 になるべきです")
}

func TestAddWithZero(t *testing.T) {
	assert := assert.New(t)

	result := Add(1, 0)
	expected := 1
	assert.Equal(expected, result, "1 + 0 は 1 になるべきです")
}
