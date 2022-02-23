package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	y := t.Year()
	i := 1996
	for {
		if i > y {
			break
		}
		fmt.Println(i)
		i++
	}
}
