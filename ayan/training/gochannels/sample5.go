package main

import (
	"fmt"
	"time"
)

func main() {

	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println(time.Now(), "Received Message:", msg)
	default:
		fmt.Println(time.Now(), "Received Message:", nil)
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println(time.Now(), "Sent Message:", msg)
	default:
		fmt.Println(time.Now(), "Sent Message:", nil)
	}

	select {
	case msg := <-messages:
		fmt.Println(time.Now(), "Received Message:", msg)
	case sig := <-signals:
		fmt.Println(time.Now(), "Received Signal:", sig)
	default:
		fmt.Println(time.Now(), nil)
	}

}
