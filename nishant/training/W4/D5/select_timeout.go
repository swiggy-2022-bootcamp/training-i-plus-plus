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

	select {
	case res := <-c1:
		fmt.Println("c1 ", res)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout on c1")
	}

	go rout(c2, 3)

	select {
	case res := <-c2:
		fmt.Println("c2 ", res)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout on c2")
	}

}
