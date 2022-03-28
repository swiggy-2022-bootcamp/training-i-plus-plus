package main

import "fmt"

// colors := make(map[string]string)
// var colors map[string]string

// colors["white"] = "#ffffff"
// fmt.Println(colors["white"])
// delete(colors, "white")

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex code for ", color, " is ", hex)
	}
}

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	printMap(colors)
	// fmt.Printf("%+v", colors)
}
