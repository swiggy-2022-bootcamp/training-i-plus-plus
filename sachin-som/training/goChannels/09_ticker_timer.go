package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1e8)  // sends time after each duration d
	boom := time.After(5e8) // sends time only once after duraction d
	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-boom:
			fmt.Println("Boom")
			return
		default:
			fmt.Println(".")
			time.Sleep(5e7)
		}
	}
}
