package main

import "fmt"

func main() {

	nameAndHobby := map[string][]string{
		"John": {"bb", "tennis", "coding"},
		"Mike": {"bb", "tennis", "code"},
	}

	nameAndHobby["Tim"] = []string{"watch", "ten", "code"}

	delete(nameAndHobby, "John")

	for k, v := range nameAndHobby {
		fmt.Println(k, v)
	}

}
