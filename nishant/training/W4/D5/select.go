package main

import (
	"fmt"
	"time"
)

func rout(ch chan int, i int) {
	time.Sleep(time.Duration(i) * time.Second)
	ch <- i
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go rout(c1, 1)
	go rout(c2, 2)

	for i := 0; i < 2; i++ {
		select {
		case msg := <-c1:
			fmt.Println("c1 :", msg)
		case msg := <-c2:
			fmt.Println("c2 :", msg)

		}
	}
}
