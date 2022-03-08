package main

import (
	"fmt"
	"sync"
)

func main() {
	msgs := make(chan string)
	var wg sync.WaitGroup

	go func() {
		msgs <- "hello"
	}()
	wg.Wait()
	go func() {
		res := <-msgs
		fmt.Println(res)
	}()
}
