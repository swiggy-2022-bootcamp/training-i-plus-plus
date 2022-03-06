package main

import "fmt"

func main() {
	library := make(map[string]string)

	library["Book-1"] = "Author-1"
	library["Book-2"] = "Author-2"
	library["Book-3"] = "Author-3"

	fmt.Println("Library :", library)
	book4, isPresent := library["Book-4"]
	fmt.Println(isPresent, book4)
}
