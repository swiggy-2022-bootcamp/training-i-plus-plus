package main

import (
	"fmt"
	"sync"
	"time"
)

func afunc(msg chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("Entered : ", i)
		msg <- i
		time.Sleep(time.Second * 1)
	}
	wg.Done()
}
func bfunc(msg chan int) {
	for i := 0; i < 10; i++ {

		fmt.Println("Received : ", <-msg)
		time.Sleep(time.Second * 2)
	}
	wg.Done()
}

var wg = &sync.WaitGroup{}

func main() {
	wg.Add(2)
	var myCh = make(chan int, 5)
	go afunc(myCh)
	go bfunc(myCh)
	wg.Wait()
}
