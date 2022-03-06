package main

import "fmt"

func main() {
	messages := make(chan string,1)
	signals := make(chan string,1)
	
	select {
	case msg := <-messages:
		fmt.Println("recived message",msg)	
	default: 
		fmt.Println("no message recieved")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message",msg)	
	default: 
		fmt.Println("no message recieved")
	}

	select {
	case msg := <-messages:
		fmt.Println("recieved message",msg)
		signals <- msg
	case sig := <-signals:
		fmt.Println("recieved signal",sig)
	default: 
		fmt.Println("no activity ")
	}

	select {
	case msg := <-messages:
		fmt.Println("recieved message",msg)
		signals <- msg
	case sig := <-signals:
		fmt.Println("recieved signal",sig)
	default: 
		fmt.Println("no activity ")
	}
}