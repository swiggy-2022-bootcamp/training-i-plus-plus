package main

import "fmt"

type animal struct {
	 name            string
	characteristics []string
}

func main() {

	animal1 := animal{
		name: "Lion",
		characteristics: []string{"Eats human",
			"Wild Animal",
			"King of the jungle",
		},
	}

	//We use dot(.) to acces each field in the struct
	fmt.Println("Animal name: ", animal1.name)
	for _, v := range animal1.characteristics {
		fmt.Printf("\t %v\n", v)
	}
}