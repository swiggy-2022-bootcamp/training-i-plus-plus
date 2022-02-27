/*
Define a 2 dimensional Point with coordinates X and Y as a struct. Do the same
for a 3 dimensional point, and a Polar point defined with its polar coordinates.
Implement a function Abs() that calculate the length of the vector represented by
a Point, and a function Scale that multiplies the coordinates of a point with a scale
factor(hint: use function Sqrt from package math).
*/

package main

import (
	"fmt"
	"math"
)

type _2DPoint struct{ x, y int }
type _3DPoint struct{ x, y, z int }
type PolarPoint struct{ r, a int }

func Abs(point *_2DPoint) (abs float32) {
	temp := point.x*point.x + point.y*point.y
	return float32(math.Sqrt(float64(temp)))
}

func main() {
	p := _2DPoint{x: 3, y: 3}
	fmt.Println(Abs(&p))
}
