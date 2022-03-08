package main

import "fmt"

type rectangle struct {
	width  int
	height int
}

func (rect rectangle) isSquare() bool {
	return rect.height == rect.width
}

func (rect *rectangle) doubleHeight() {
	rect.height *= 2
}

func main() {
	var rect = rectangle{width: 10, height: 10}
	fmt.Println(rect.isSquare())
	rect.doubleHeight()
	fmt.Println(rect.isSquare())
}
