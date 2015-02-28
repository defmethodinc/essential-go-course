package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNumberLine(t *testing.T) {
	assert.Equal(t, "0 1 2 3 4 5 6 7 8 9", GetLine(0))
	assert.Equal(t, "10 11 12 13 14 15 16 17 18 19", GetLine(10))
	assert.Equal(t, "20 21 22 23 24 25 26 27 28 29", GetLine(20))
}
