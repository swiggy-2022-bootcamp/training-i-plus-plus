package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	y := t.Year()
	for i := 1996; i <= y; i++ {
		fmt.Println(i)
	}
}
