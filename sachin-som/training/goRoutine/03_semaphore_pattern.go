package main

import (
	"fmt"
	"time"
)

type Empty interface{}
type semaphore chan Empty

// Release n resources
func (s semaphore) signal(n int) {
	for i := 0; i < n; i++ {
		s <- i
	}
}

// Acuire n resources
func (s semaphore) wait(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

var (
	n   = 9
	sem = make(semaphore, n)
)

func main() {
	dataChan := make(chan int, n)
	go producer(n, dataChan)
	go consumer(n, dataChan)
	// time.Sleep(10 * time.Second)
	sem.wait(n)
}

func producer(n int, ch chan int) {
	for i := 0; i < n; i++ {
		time.Sleep(1 * time.Second)
		ch <- i * i
	}
}

func consumer(n int, ch chan int) {
	for i := 0; i < n; i++ {
		fmt.Printf("Received Data: %d\n", <-ch)
	}
	sem.signal(n)
}
