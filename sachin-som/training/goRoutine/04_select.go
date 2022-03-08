package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)
	time.Sleep(time.Second / 100)
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Data received from ch1: %v\n", v)
		case v2 := <-ch2:
			fmt.Printf("Data received from ch2: %v\n", v2)
		default:
			fmt.Println("Didn't received any data.")
		}
	}
}
