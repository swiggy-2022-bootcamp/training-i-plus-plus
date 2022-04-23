package main

import (
	"fmt"
	"math/rand"
	"time"
)

var messages = make(chan string)

func sendMessage(msg string) {
	seconds := rand.Intn(10)
	fmt.Println(msg, " : this will take ", seconds, " seconds")
	time.Sleep(time.Duration(seconds) * time.Second)
	messages <- msg
}

func worker(done chan bool) {
	fmt.Print("It's working !!!")
	time.Sleep(time.Second)
	fmt.Println("DONE !")

	done <- true
}

func main() {

	go sendMessage("Tik..  ")
	go sendMessage("Tok..  ")

	messageRecieved := <-messages
	fmt.Println(messageRecieved)

	messages := make(chan string, 2)

	messages <- "Buffered"
	messages <- "Channel in Go"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	done := make(chan bool, 1)
	go worker(done)

	<-done

	messages1 := make(chan string, 1)
	signals := make(chan bool, 1)

	select {
	case msg := <-messages1:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages1 <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages1:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

	close(done)
	close(messages)
	close(messages1)

}
