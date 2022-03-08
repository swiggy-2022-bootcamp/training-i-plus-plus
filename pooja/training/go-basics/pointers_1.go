package main

import "fmt"

type Creature struct {
	Species string
}

func main() {
	var creature *Creature
	creature = &Creature{Species: "shark"}
	fmt.Printf("1) %+v\n", creature)
	changeCreature(creature)
	fmt.Printf("3) %+v\n", creature)
}

func changeCreature(creature *Creature) {
	if creature == nil {
		fmt.Println("creature is nil")
		return
	}
	creature.Species = "jellyfish"
	fmt.Printf("2) %+v\n", creature)
}
