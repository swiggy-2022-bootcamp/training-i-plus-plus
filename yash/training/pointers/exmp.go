package main

import "fmt"

type Creature struct {
	species string
}

func (c *Creature) Reset() {
	c.species = ""
}
func main() {
	var creature *Creature = &Creature{species: "dog"}
	fmt.Printf("1) %+v\n", creature)
	changeCreature(creature)
	fmt.Printf("3) %+v\n", creature)

	fmt.Println("\nReset method")
	var c Creature = Creature{species: "cat"}
	fmt.Printf("1) %+v\n", c)
	c.Reset()
	fmt.Printf("2) %+v\n", c)

}
func changeCreature(creature *Creature) {
	if creature == nil {
		fmt.Println("creature is nil")
		return
	}
	creature.species = "human"
	fmt.Printf("2) %+v\n", creature)
}
