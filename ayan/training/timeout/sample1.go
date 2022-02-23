package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "one"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "two"
	}()

	select {
	case msg2 := <-ch2:
		fmt.Println(msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
