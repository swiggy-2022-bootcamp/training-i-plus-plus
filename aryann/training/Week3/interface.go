package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func calArea(s shape) {

	switch s.(type) {
	case square:
		fmt.Println("Area of square: ", s.area())
	case circle:
		fmt.Println("Area of circle: ", s.area())
	default:
		fmt.Println("Unknown shape")
	}
}

func main() {

	c := circle{
		radius: 5,
	}

	s := square{
		length: 5,
	}

	calArea(c)
	calArea(s)

	fmt.Println("Area of circle: ", c.area())
	fmt.Println("Area of square: ", s.area())

}
