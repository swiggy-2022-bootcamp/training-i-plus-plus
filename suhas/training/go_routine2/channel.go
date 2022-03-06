package main

import "fmt"
//accepts a channel for sending values
func ping(pings chan<- string, msg string) {
	pings<-msg
}

//accepts one channel for recevies (pings) and a second for sends(pongs)
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <-msg
}

func numbers() {
	for i:=1; i<=50; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ",i)
		ping(pings, "passed message")
		pong(pings,pongs)
	}
}

func alphabets() {
	for i:='a'; i<='z'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ",i)
		ping(pings, "passed message")
		pong(pings,pongs)
	}
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	go numbers()
	go alphabets()

	fmt.Println(<-pongs)
}