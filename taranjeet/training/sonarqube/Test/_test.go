package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDemo(t *testing.T) {
	msg1 := "hello"
	msg2 := "hello"

	assert.Equal(t, msg1, msg2, "two msg should be same")

}
