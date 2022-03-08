package main

import "fmt"

func main() {

	nameAndHobby := map[string][]string{
		"Steven": {"Basket Ball", "Table tennis", "Coding"},
		"John":   {"Sleeping", "Netflix", "Swimming"},
	}

	nameAndHobby["Timi"] = []string{"Watching cartoon", "Anime", "Studying"}

	delete(nameAndHobby,"Steven")
	for i, v := range nameAndHobby {
		fmt.Printf("%v likes \n",i)
		for j, value := range v{
			fmt.Printf("\t%v \t%v\n", j, value)
		}
	}
}