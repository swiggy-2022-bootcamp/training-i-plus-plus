package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {

	msg1 := "Hello"
	msg2 := "Hello"

	assert.Equal(t, msg1, msg2)
}
