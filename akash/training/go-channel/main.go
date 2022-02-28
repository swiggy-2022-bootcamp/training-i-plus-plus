package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hey")

	myChannel := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	fmt.Println(myChannel)
	fmt.Printf("%T\n", myChannel)

	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-myChannel)
		wg.Done()
	}(myChannel, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		myChannel <- 5
		wg.Done()
	}(myChannel, wg)

	wg.Wait()
}
