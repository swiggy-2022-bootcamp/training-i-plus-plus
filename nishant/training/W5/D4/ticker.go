package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("ticker stopped")
				return
			case t := <-ticker.C:
				fmt.Println("tick ", t)
			}
		}
	}()

	time.Sleep(2000 * time.Millisecond)
	ticker.Stop()
	done <- true
}
