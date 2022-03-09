package main

import (
	"fmt"
	"time"
)

func boring(str string, channel chan int) {
	for i := 0; ; i++ {
		channel <- i + 1
		time.Sleep(1 * time.Second)
	}
}

func main() {

	start := time.Now()

	c7 := make(chan int)
	go boring("boring", c7)
	for i := 0; i < 5; i++ {
		fmt.Printf("Boring %dth times\n", <-c7)
	}

	t := time.Now()
	elapsedTime := t.Sub(start)
	fmt.Printf("Starting for fourth - %v Ending for fourth - %v, Time taken - %v \n", start, t, elapsedTime)

	fmt.Println("Total Time - ", elapsedTime)

}
