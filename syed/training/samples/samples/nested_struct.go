package main

import "fmt"

type animal struct {
	name            string
	characteristics []string
}

//A herbivore is an animal, so it can have the animal struct as a field
type herbivore struct {
	animal
	eatHuman bool
}

func main() {

	herb := herbivore{
		animal: animal{
			name: "Goat",
			characteristics: []string{"Lacks sense",
				"Lazy",
				"Eat grass",
			},
		},
		eatHuman: false, //maybe
	}

	//We use dot(.) to acces each field in the struct
	fmt.Println("Animal name:", herb.animal.name)
	fmt.Println("Eats human? ", herb.eatHuman)
	fmt.Println("Characteristics:")
	for _, v := range herb.animal.characteristics {
		fmt.Printf("\t %v\n", v)
	}
}