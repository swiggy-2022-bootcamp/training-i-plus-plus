package main

import "fmt"

func main() {
	colors := map[string]string {
		"red": "#ff0000",
		"green": "#4bf745",
		"white" : "#ffffff",
	}

	// colors := map[string]string {	
	// }

	// var colors map[string]string

	//To create empty map, we can also use make function
	// colors := make(map[string]string)

	//to add key
	// colors["white"] = "#ffffff"

	//to delete key
	// delete(colors, "white")

	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}