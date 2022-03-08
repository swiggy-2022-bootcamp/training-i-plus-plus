package main

import "fmt"

type Creature struct {
	Species string
}

//does not change anything
func (c Creature) Reset() {
	c.Species = ""
}

func (c *Creature) ResetViaPointer() {
	c.Species = ""
}
func main() {
	var creature = &Creature{Species: "shark"}
	fmt.Printf("1) %+v\n", creature)
	creature.Reset()
	fmt.Printf("2) %+v\n", creature)
	creature.ResetViaPointer()
	fmt.Printf("3) %+v\n", creature)
}
