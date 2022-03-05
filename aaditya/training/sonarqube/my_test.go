package Test


import (

	"testing"


	"github.com/stretchr/testify/assert"

)


func TestDemo(t *testing.T) {

	one := "one"

	two := "one"


	assert.Equal(t, one, two, "the two variables should be the same value")

}

func TestDemo1(t *testing.T) {

	one := "one"

	two := "two"


	assert.Equal(t, one, two, "the two variables should be the same value")

}

