package main

import . "fmt"

type Rectangle struct {
	Name          string
	Width, Height float64
}

func main() {
	var a Rectangle
	var b = Rectangle{"I'm b.", 10, 20}
	var c = Rectangle{Height: 12, Width: 14}

	Println(a)
	Println(b)
	Println(c)
}