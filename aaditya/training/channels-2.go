package main

import (
	"fmt"
	"time"
)

func ping(pings chan<- string, msg string){
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string){
	msg := <-pings
	pongs <- msg
}

func main(){
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	go ping(pings, "passed msg")
	go pong(pings,pongs)

	fmt.Println(<-pongs)

	time.Sleep(10 * time.Second)
}