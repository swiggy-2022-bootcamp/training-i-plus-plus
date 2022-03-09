package main

import (
	"fmt"
	"time"
)

func gorout(count int, msg string) {
	for i := 0; i < count; i++ {
		fmt.Println(msg)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	count := 10

	go gorout(count, "Inside Go Routine")

	gorout(count, "From Main")
}
