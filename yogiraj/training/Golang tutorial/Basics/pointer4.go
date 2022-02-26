package main

import "fmt"

type Creature struct{
	species string
}

func main(){
	var creature Creature = Creature{"shark"}
	
	fmt.Printf("1) %+v\n", creature)
	changeCreature(&creature)
	fmt.Printf("3) %+v\n", creature)
}

func changeCreature(creature *Creature){
	
	creature.species="jellyfish"
	fmt.Printf("2) %+v\n", creature)
}