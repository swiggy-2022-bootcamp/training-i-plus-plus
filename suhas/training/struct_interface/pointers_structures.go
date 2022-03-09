package main 

import "fmt"

type Creature struct {
	Species string	
}

func main() {

	var creature Creature = Creature{Species:"Shark"}

	fmt.Print("f1) %+v\n",creature)
	changeCreature(&creature)
	fmt.Printf("3) %+v\n",creature)
}


func changeCreature(creature *Creature) {
	creature.Species = "Jellyfish"
	fmt.Printf("2) %+v\n",creature)
}