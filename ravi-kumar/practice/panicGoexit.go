package main

import (
	"runtime"
)

//panic and Goexit are independent of each other. So one isn't going to replace another, unlike multiple panics do (see panicRecover.go).
//but go compiler doesnt implement it the right way - overshadowing does happen in standard go compilers. This issue is to be resolved in future compilers.
func main() {
	// The Goexit signal shadows the "1"
	// panic now, but it should not.
	defer runtime.Goexit()
	panic(1)
}
