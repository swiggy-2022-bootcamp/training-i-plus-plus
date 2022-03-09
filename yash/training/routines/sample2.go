package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int) {
	fmt.Printf("Task %d is running\n", id)
	time.Sleep(time.Second * 2)
	fmt.Printf("Task %d is done\n", id)
}
func main() {
	var waitg sync.WaitGroup
	for i := 0; i < 10; i++ {
		waitg.Add(1)
		i := i
		go func() {
			defer waitg.Done()
			task(i)
		}()
		waitg.Wait()
	}
}
