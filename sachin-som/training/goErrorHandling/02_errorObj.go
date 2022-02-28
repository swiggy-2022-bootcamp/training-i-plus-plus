package main

import (
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("math: square root of negative number %g", f) // would return an error object
	}
	return math.Sqrt(f), nil
}
func main() {
	sqrt, err := Sqrt(-9)
	if err == nil {
		fmt.Println(sqrt)
	} else {
		fmt.Println(err)
	}
}
