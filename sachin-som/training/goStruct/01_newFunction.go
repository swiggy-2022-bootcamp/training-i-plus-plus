package main

import "fmt"

// Struct Example
type struct1 struct {
	i1  int
	f1  float64
	str string
}

func main() {

	// New function returns a pointer to
	// to the allocated memory block of
	// the struct
	var s *struct1
	s = new(struct1) // takes a type and returns zero value pointer of the memory
	fmt.Println(s)
	s.f1 = 9.4
	s.i1 = 89
	s.str = "sachin"
	fmt.Println(s)

	// Shorter way
	s = &struct1{90, 90.0, "sachin"}

	fmt.Println(s)
}
