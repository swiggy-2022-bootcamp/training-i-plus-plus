package main

import (
	"fmt"
	"sync"
)

func f(s string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	fmt.Println(s)
}

func main() {
	var waitGroup sync.WaitGroup 
	waitGroup.Add(1)
	go f("i'm async!", &waitGroup)

	waitGroup.Add(1)
	go func(s string) {
		defer waitGroup.Done()
		fmt.Println(s)
	}("i'm async too!")

	waitGroup.Wait()

	fmt.Println("done!")
}

// WaitGroup is a struct defined in the sync package. It maintains a counter to wait for a collection of goroutines to finish. In the above program
// waitGroup.Add(1) adds a delta to the WaitGroup counter
// waitGroup.Done() decreases the WaitGroup counter by one
// waitGroup.Wait() blocks until the WaitGroup counter is zero
// Now, the quantity of output strings is always consistent (the ordering may be different for each invocation)