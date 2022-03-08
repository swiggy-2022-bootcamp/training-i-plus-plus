package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(4 * time.Second)
		c1 <- "result 1"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c2:
			fmt.Println("recieved ", msg1)
		case msg2 := <-c1:
			fmt.Println("recieved ", msg2)
		}
	}
}
