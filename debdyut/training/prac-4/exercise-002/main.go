package main

import "fmt"

func main() {
	a := 42 == 42
	b := 43 >= 42
	c := 43 <= 42
	d := 43 < 42
	e := 43 > 42
	fmt.Println(a, b, c, d, e)
}
