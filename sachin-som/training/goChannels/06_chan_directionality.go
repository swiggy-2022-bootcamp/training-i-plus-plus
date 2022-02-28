/*
var send_only chan<- int // channel can only receive data
var recv_only <-chan int // channel can only send data
*/

package main

import "fmt"

func main() {
	ch := make(chan int)
	done := make(chan bool)

	go func(ch chan<- int) {
		ch <- 333
	}(ch)

	go func(ch <-chan int, done chan<- bool) {
		fmt.Println(<-ch)
		done <- true
	}(ch, done)
	<-done
}
