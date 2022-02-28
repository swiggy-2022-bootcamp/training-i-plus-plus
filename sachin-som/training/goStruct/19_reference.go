package main

import "fmt"

type Creature struct {
	Species string
}

func ChangeSpecies(creature *Creature, sp string) {
	creature.Species = sp
}

func main() {
	s := Creature{Species: "Human"}
	ChangeSpecies(&s, "Animal")
	fmt.Println(s)
}
