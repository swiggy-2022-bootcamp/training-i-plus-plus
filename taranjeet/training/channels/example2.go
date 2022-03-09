package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	c3 := make(chan string)
	c4 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c3 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c4 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c3:
			fmt.Println("received", msg1)
		case msg1 := <-c4:
			fmt.Println("received", msg1)

		}

	}
	t := time.Now()
	elapsedTime := t.Sub(start)
	fmt.Printf("Starting for second - %v Ending for second - %v, Time taken - %v \n", start, t, elapsedTime)

}
