package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		chan2 <- "two"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		chan1 <- "one"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			fmt.Println("received ", msg1)
		case msg2 := <-chan2:
			fmt.Println("received ", msg2)
		}
	}
}
