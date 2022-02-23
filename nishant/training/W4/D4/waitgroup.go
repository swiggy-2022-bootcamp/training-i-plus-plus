package main

import (
	"fmt"
	"sync"
	"time"
)

func sleeper(i int) {
	fmt.Println("start : ", i)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("end : ", i)
}

func main() {

	var waitg sync.WaitGroup

	for i := 0; i < 10; i++ {
		waitg.Add(1)
		k := i
		go func() {
			defer waitg.Done()//3
			defer // 2
			defer // 1
			sleeper(k)
		}()
	}

	//waitg.Wait()
}
