package main

import "fmt"

type Person struct {
	name string
}

func main() {
	p1 := Person{"nishant"}
	p1.hi()
	p1.hi()

	fmt.Println(multi())
}

func (p *Person) hi() {
	fmt.Printf(" hello from %s \n", p.name)
	p.name = "asdfdsfa"
}

func multi() (a, b string) {
	a = "a"
	b = "b"
	return
}
