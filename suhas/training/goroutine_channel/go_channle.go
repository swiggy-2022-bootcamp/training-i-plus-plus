package main

import (
	"fmt"
	"time"
)
func main() {
	msgs := make(chan string)

	go func() {
		msgs <- "hello"
	}()

	go func() {
		response := <- msgs
		fmt.Println(response)
	}()

	time.Sleep(time.Second)
}