/*
Define a struct Rectangle with int properties length and width. Give this type methods Area() and
Perimeter() and test it out
*/

package main

import "fmt"

type Rectangle struct{ length, width int }

func (r *Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func (r *Rectangle) Area() int {
	return (r.length * r.width)
}

func main() {
	rect1 := &Rectangle{length: 4, width: 4}
	fmt.Println(rect1.Area())
}
