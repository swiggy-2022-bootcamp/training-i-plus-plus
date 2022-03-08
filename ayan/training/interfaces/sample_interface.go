package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
	// disp()
}

func (s square) area() float64 {
	return s.side * s.side
}

func (s square) disp() {
	fmt.Println("I am a square with side", s.side)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func calcArea(s shape) {

	switch s.(type) {
	case square:
		// s.disp()
		fmt.Println("Area of square : ", s.area())
	case circle:
		fmt.Println("Area of circle : ", s.area())
	}
}

func main() {
	c := circle{2.5}
	s := square{2.5}

	calcArea(c)
	calcArea(s)
}
