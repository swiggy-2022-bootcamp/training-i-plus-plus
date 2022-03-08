package main

import "fmt"

// pings - write-only channel
func ping(pings chan<- string, msg string) {

	pings <- msg
}

// pings - read-only channel; pongs - write-only channel
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
