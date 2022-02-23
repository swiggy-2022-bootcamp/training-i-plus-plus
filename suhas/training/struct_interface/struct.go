package main 

import "fmt"

type rectangle struct {
	length int
	bradth int
	color string
}


func main() {
	fmt.Println(rectangle{10,20,"red"})
}