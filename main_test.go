package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert := assert.New(t)
	a := &Adder{}

	a.num1 = 1
	a.num2 = 2

	assert.Equal(3, a.Calculate(), "2 + 1 should be 3")
}
