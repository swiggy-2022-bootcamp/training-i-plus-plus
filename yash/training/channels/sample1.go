// Channel are pipes that connect concurrent goroutines.

package main

import (
	"fmt"
	"time"
)

func main() {
	msgs := make(chan string)
	go func() {
		msgs <- "Hello" // sending message to channel
	}() // this is a go routine
	time.Sleep(time.Second)

	res := <-msgs // receiving message from the channel
	fmt.Println(res)

}
