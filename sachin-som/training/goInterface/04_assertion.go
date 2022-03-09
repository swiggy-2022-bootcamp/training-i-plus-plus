/*


 */

package main

import (
	"fmt"
	"math"
)

type Circle2 struct {
	radius float32
}

type Rectange2 struct {
	length float32
	width  float32
}

type Shapper2 interface {
	Area() float32
}

func (s *Rectange2) Area() float32 {
	return s.length * s.width
}

func (c *Circle2) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func main() {

	var shapper Shapper2

	shapper = &Rectange2{length: 2, width: 4}
	// shapper = &Circle2{radius: 3.0}

	// Checked type assertion
	// if t, ok := shapper.(*Rectange2); ok {
	// 	fmt.Printf("Base type of interface is: %T\n", t)
	// }

	// type-switch
	switch shapper.(type) {
	case *Rectange2:
		fmt.Println("rectangle type")
	case *Circle2:
		fmt.Println("Circle type")
	case nil:
		fmt.Println("Nothing to check")
	default:
		fmt.Println("Other type")
	}

	Classifier(13, -14.3, "BELGIUM", complex(1, 2), nil, false)
}

func Classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("param #%d is a bool\n", i)
		case float64:
			fmt.Printf("param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("param #%d is an int\n", i)
		case nil:
			fmt.Printf("param #%d is nil\n", i)
		case string:
			fmt.Printf("param #%d is a string\n", i)
		default:
			fmt.Printf("param #%d’s type is unknown\n", i)
		}
	}
}
