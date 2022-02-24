package main

import (
	"fmt"
)

type Shapper interface {
	Area() float64
}

type Square struct {
	side int
}

func (s *Square) Area() float64 {
	return float64(s.side) * float64(s.side)
}

type Rectange struct {
	length, width int
}

func (r *Rectange) Area() float64 {
	return float64(r.length) * float64(r.width)
}

func main() {
	shapes := []Shapper{&Square{side: 4}, &Rectange{length: 4, width: 5}}
	for _, shape := range shapes {
		fmt.Println(shape.Area())
	}
}
