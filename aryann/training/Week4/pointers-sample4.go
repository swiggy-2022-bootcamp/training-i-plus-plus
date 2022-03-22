package main

import "fmt"

type Creature struct {
	Species string
}

func main() {

	var creature Creature = Creature{Species: "shark"}

	fmt.Printf("1) %+V\n", creature)

	changeCreature(creature)

	fmt.Printf("3) %+V\n", creature)
}

func changeCreature(creature Creature) {

	creature.Species = "whale"

	fmt.Printf("2) %+V\n", creature)
}
