package main

import "fmt"

type T struct {
	V  int
	tt *T
}

func (t *T) hello() string {
	return "world"
}

type A struct {
	a int
}

func main() {
	var t *T = nil
	fmt.Println(t)            //nil
	fmt.Println(t.tt.hello()) //panic
}
