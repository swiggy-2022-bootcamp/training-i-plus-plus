package main

import "fmt"

const (
	y1 = 2022 - iota
	y2 = 2022 - iota
	y3 = 2022 - iota
	y4 = 2022 - iota
)

func main() {
	fmt.Println(y1, y2, y3, y4)
}
