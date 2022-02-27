package main

import "fmt"

type Creature struct {
	Species string
}
func main(){
	var creature *Creature
	creature = &Creature{Species: "shark"}
	fmt.Printf("1) %+v\n",*creature)
	changeCreature(creature)
	fmt.Printf("2) %+v\n",*creature)
}

func changeCreature(creature *Creature){
	if creature == nil {
		fmt.Println("creature doesn't hold any value")
		return
	}
	creature.Species = "jellyfish"
	fmt.Printf("\nCreatue -> ", creature)
	fmt.Printf("\n&Creature -> ", &creature)
	fmt.Printf("\n*Creature -> ", *creature)
	return
}