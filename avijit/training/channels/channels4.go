package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer((2 * time.Second))
	<-timer1.C
	fmt.Println(" Time 1 Fired ")

	timer2 := time.NewTimer((time.Second))
	go func() {
		<-timer2.C
		fmt.Println((" Timer 2 Fired"))
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println(" Timer 2 Stopped")
	}

	time.Sleep(2 * time.Second)

}
