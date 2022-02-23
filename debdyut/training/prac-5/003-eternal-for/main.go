package main

import "fmt"

func main() {
	x := 0
	// similar to while(true) or for(;;)
	for {
		x++
		if x >= 25 {
			break
		}
		if x%2 == 0 {
			continue
		}
		fmt.Println(x)

	}
}
