/*
Define an interface Simpler with methods Get() which returns an integer, and Set() which has an
integer as parameter. Make a struct type Simple which implements this interface.
Then define a function which takes a parameter of the type Simpler and calls both methods upon
it. Call this function from main to see if it all works correctly
*/

package main

import "fmt"

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	val int
}

func (s *Simple) Get() int {
	return s.val
}

func (s *Simple) Set(val int) {
	s.val = val
}

func example(s Simpler) {
	s.Set(20)
	fmt.Println(s.Get())
}
func main() {
	s := &Simple{}
	example(s)
}
