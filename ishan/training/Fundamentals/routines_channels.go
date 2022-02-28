package main

import (
	"fmt"
	"time"
)

func receiver(ch chan int) {
	fmt.Println(<-ch)
}

func main() {
	fmt.Println("Channels and routines")

	ch := make(chan int)

	go receiver(ch)

	ch <- 42

	time.Sleep(1000 * 3600)
}
