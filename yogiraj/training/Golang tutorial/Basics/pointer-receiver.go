package main

import "fmt"

type Creature struct{
	species string
}

func (c *Creature) Reset(){
	c.species=""
}

func main(){
	var creature Creature = Creature{"shark"}
	
	fmt.Printf("1) %+v\n", creature)
	creature.Reset()
	fmt.Printf("3) %+v\n", creature)
}

