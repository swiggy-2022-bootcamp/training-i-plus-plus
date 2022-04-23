package main

import (
	"fmt"
	. "sonarcubeApp/math"
)

func main() {
	var res int
	res = Add(1, 2)
	fmt.Println(res)

	res = Sub(1, 2)
	fmt.Println(res)

	res = Mul(1, 2)
	fmt.Println(res)

	res = Div(1, 2)
	fmt.Println(res)
}
