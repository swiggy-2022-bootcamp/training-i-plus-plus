package main

import (
	"fmt"
	"time"
)

func main() {

	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		channel1 <- "ping"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- "pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-channel1:
			fmt.Println("message recieved : ", msg1)

		case msg2 := <-channel2:
			fmt.Println("message recieved : ", msg2)
		}
	}

}
