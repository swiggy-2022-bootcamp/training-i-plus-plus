package main

import (
	"fmt"
	"math/rand"
	"time"
)

var messages = make(chan string)

func sendMessage(msg string) {
	seconds := rand.Intn(10)
	fmt.Println(msg, " : will take ", seconds, " seconds")
	time.Sleep(time.Duration(seconds) * time.Second)
	messages <- msg
}

func main() {

	go sendMessage("Ping")
	go sendMessage("Pong")

	messageRecieved := <-messages
	fmt.Println(messageRecieved)
}
