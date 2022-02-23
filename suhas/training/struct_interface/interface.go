package main 


import (
	"fmt"
	"math"
)

type square struct {
	lenght float64
}

type circle struct {
	radius float64
}


type rectangle struct {
	length float64
	breadth float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.lenght * s.lenght
}

func (c circle) area() float64 {
	return math.Pi * c. radius * c.radius
}

func calArea(s shape) {
	switch s.(type) {
		case square:
			fmt.Println("This is the area of the sqaure : ",s.area())
		case circle:
			fmt.Println("This is the area of the circle : ",s.area())
	}
}

func main() {
	c := circle{
		radius: 2.5,
	}
	s := square {
		lenght: 1.5,
	}
	r := rectangle {
		lenght : 1.15,
		breadth: 2.25,
	}

	calArea(c)
	calArea(s)
	//calArea(r) missing area method
}