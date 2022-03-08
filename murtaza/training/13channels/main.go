package main

import (
	"fmt"
	"sync"
	"time"
)

func declareChannelDemo() {
	var mychannel1 chan int
	fmt.Println(mychannel1)

	mychannel2 := make(chan int)
	fmt.Println(mychannel2)
}

func publisher(wg *sync.WaitGroup, channel chan int, limit int) {
	defer wg.Done()
	for i := 1; i <= limit; i++ {
		channel <- i
		time.Sleep(time.Second)
	}
}

func consumer(wg *sync.WaitGroup, channel chan int, limit int) {
	if wg != nil {
		defer wg.Done()
	}

	for i := 1; i <= limit; i++ {
		fmt.Println(<-channel)
	}
}

func demoWithStandardChannel() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	mychannel := make(chan int)

	defer close(mychannel)
	go publisher(wg, mychannel, 5)
	go consumer(wg, mychannel, 5)

	wg.Wait()
}

func demoWithBufferedChannel() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	capacity := 2
	myBufferedChannel := make(chan int, capacity)

	defer close(mychannel)
	go publisher(wg, myBufferedChannel, capacity)
	wg.Wait()

	consumer(nil, myBufferedChannel, capacity)
}

func main() {
	demoWithBufferedChannel()
}
