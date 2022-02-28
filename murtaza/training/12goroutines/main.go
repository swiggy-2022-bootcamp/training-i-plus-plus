package main

import (
	"fmt"
	"sync"
	"time"
)

func loopAndPrint(wg *sync.WaitGroup, text string, limit int) {
	defer wg.Done()
	for i := 0; i < limit; i++ {
		fmt.Println(text, " : ", i)
		time.Sleep(time.Second)
	}
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go loopAndPrint(wg, "Hello there", 4)
	go loopAndPrint(wg, "Heyaa", 10)
	wg.Wait()
}
