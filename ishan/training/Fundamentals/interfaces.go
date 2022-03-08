package main

import (
	"fmt"
	"math"
)

type rect struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

type geometry interface {
	perimeter() float64
}

func measure(g geometry) { fmt.Println(g.perimeter()) }

func (r rect) perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	r := rect{
		length: 2,
		width:  4,
	}
	c := circle{
		radius: 2,
	}
	measure(r)
	measure(c)
}
