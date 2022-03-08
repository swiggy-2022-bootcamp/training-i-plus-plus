package main

import (
	"fmt"
	"sync/atomic"

	"sync"
)

func main() {

	var race uint64
	var atm uint64
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for k := 0; k < 500; k++ {
				atomic.AddUint64(&atm, 1)
				race++
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("race : ", race)
	fmt.Println("atm : ", atm)

}
