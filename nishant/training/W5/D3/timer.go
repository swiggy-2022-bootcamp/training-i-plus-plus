package main

import (
	"fmt"
	"time"
)

func main() {

	t1 := time.NewTimer(1 * time.Second)

	<-t1.C

	fmt.Println("timer 1 done")

	t2 := time.NewTimer(3 * time.Second)

	go func() {
		<-t2.C
		fmt.Println("timer 2 done")
	}()

	t2.Stop()
	fmt.Println("t2 stopped")
}
