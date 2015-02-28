package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDouble(t *testing.T) {
	assert.Equal(t, 10, Double(5))
}

func TestParseName(t *testing.T) {
	first, last := ParseName("Joe Leo")
	assert.Equal(t, "Joe", first)
	assert.Equal(t, "Leo", last)
}
