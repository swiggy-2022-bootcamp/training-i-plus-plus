package main

import "fmt"

type Point struct{ x, y int }

func main() {

	// using new method
	var s *Point
	s = new(Point)
	s.x = 2
	s.y = 3
	fmt.Println(*s)

	// Using struct literal
	var s1 *Point
	s1 = &Point{x: 20, y: 10}
	fmt.Println(*s1)
}
