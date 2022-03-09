package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	c5 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c5 <- "res 1"
	}()
	select {
	case res := <-c5:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
		t := time.Now()
		elapsedTime := t.Sub(start)
		fmt.Println(elapsedTime)

	}

	c6 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c6 <- "res 2"
	}()
	select {
	case res := <-c6:
		fmt.Println(res)
		t := time.Now()
		elapsedTime := t.Sub(start)
		fmt.Println(elapsedTime)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")

	}

	t := time.Now()
	elapsedTime := t.Sub(start)
	fmt.Printf("Starting for third - %v Ending for third - %v, Time taken - %v \n", start, t, elapsedTime)

}
