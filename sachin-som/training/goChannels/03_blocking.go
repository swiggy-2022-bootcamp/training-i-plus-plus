package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	// Receives data till 1 sec and then stops
	for start := time.Now(); time.Since(start) < time.Second; {
		fmt.Println(<-ch)
	}

	// program ends as ch blocks
}
