package main

import (
	"errors"
	"fmt"
)

/*
error interface
---------------
type error interface {
	Error() string
}


errorString
type errorString struct {
	s string
}

func New(text string) error {
	return &errorString{s: text}
}
*/
var err error = errors.New("Example error") // Returns address of an errorString struct which implements 'error' interface
func main() {
	fmt.Println(err.Error())
}
