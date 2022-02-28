package main

import "fmt"

type Person2 struct {
	name string
	age  int
}

type Student2 struct {
	Person2
	std int
}

func (p *Person2) printValues() {
	fmt.Println(p)
}

func (s *Student2) printValues() {
	fmt.Println(s)
}

func main() {
	p := &Person2{name: "sachin", age: 12}
	s := &Student2{Person2: *p, std: 10}

	p.printValues()
	s.printValues()
}
