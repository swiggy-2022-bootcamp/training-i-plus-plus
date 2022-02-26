/*
instead of passing a channel as a parameter to a goroutine,
let the function make the channel and return it (so it plays
the role of a factory); inside the function a lambda function
is called as a goroutine.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	suck(pump())
	time.Sleep(time.Second)
}

func pump() chan int {

	// create a channel
	ch := make(chan int)

	// Do stuff, inside a go routine
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	// return the channel
	return ch
}

func suck(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}
