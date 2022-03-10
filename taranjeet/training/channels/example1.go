package main

import (
	"fmt"
	"sync"
	"time"
)

func sendToChan1(wg *sync.WaitGroup, chan1 chan string) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	chan1 <- "hello"
	chan1 <- "world"

}

func sendToChan2(wg *sync.WaitGroup, chan2 chan int) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	chan2 <- 5

}

func main() {

	start := time.Now()

	// example 1

	chan1 := make(chan string)
	chan2 := make(chan int)

	wg := sync.WaitGroup{}

	wg.Add(2)

	go sendToChan1(&wg, chan1)

	go sendToChan2(&wg, chan2)

	fmt.Println(<-chan1)
	fmt.Println(<-chan2)
	fmt.Println(<-chan1)

	t := time.Now()
	elapsedTime := t.Sub(start)
	fmt.Printf("Starting for first - %v Ending for first - %v, Time taken - %v \n", start, t, elapsedTime)

	wg.Wait()

}
