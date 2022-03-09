package main

import "fmt"
// accepts a channel for sending values
func ping(pings chan<- string, msg string) {
    pings <- msg
}
// accepts one channel for receives (pings) and a second for sends (pongs)
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings //read
    pongs <- msg   //write
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)

    ping(pings, "passed message")
    pong(pings, pongs)

    fmt.Println(<-pongs)
}