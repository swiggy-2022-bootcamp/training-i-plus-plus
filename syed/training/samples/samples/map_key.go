package main

import "fmt"

func main() {

	//map with a key of string and value of a slice of string:

	nameAndHobby := map[string][]string{
		"Steven": []string{"Basket Ball", "Table Tennis", "Coding"},
		"Nnamdi": []string{"Sleeping", "Watching Movie", "Eating"},
	}
	//We can add someone else with their hobby:
	nameAndHobby["Timi"] = []string{"Watching Cartoon", 
					"Dreaming", 
					"Laughing", 
					"Lazing around",
				 }

	//We can delete from the map:
	delete(nameAndHobby, "Steven")

	for i, v := range nameAndHobby {
		fmt.Printf("%v likes \n", i)
		for j, value := range v {
			fmt.Printf("\t%v %v\n", j, value)
		}
	}
}