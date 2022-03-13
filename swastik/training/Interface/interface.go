package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type rect struct {
	length float64
	width float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64{
	return  (math.Pi * c.radius * c.radius)
}

func (r rect) area() float64{
	return r.length * r.width
}

func main() {
	c1 := circle{2}
	r1 := rect{4,5}
	shapes := []shape{c1,r1}

	for index ,shape := range shapes{
		fmt.Println(index, shape.area())
	}
}
