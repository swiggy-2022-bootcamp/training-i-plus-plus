package main

import (
	"fmt"
)

func get(ch chan int, done chan bool) {
	fmt.Println(<-ch)
	done <- true
}

func send(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
}

func main() {
	// Channel will able to hold 5 values and will not block untill 5 values
	ch := make(chan int, 5)
	done := make(chan bool)
	go send(ch)
	go get(ch, done)
	// time.Sleep(2 * time.Second)
	<-done
}
