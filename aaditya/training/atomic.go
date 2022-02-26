package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops uint64
	var wg sync.WaitGroup
	for i := 0; i< 5; i++ {
		wg.Add(1)
		go func ()  {
			for c:=0; c < 10; c++ {
				atomic.AddUint64(&ops,1)
				fmt.Println(c, "->", ops)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("ops: ", ops)
}