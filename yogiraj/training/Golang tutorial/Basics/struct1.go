package main

import "fmt"

type animal struct {
	name            string
	characteristics []string
}

func main() {
	animal1 := animal{
		name:            "Lion",
		characteristics: []string{"Eats animal", "Wild animal", "King of jungle"},
	}

	fmt.Println("Animal name: ", animal1.name)
	for _,v := range animal1.characteristics{
		fmt.Printf("\t %v\n",v)
	}
}