package main

import "fmt"

func main() {
	var (
		firstname, lastname, s string
		age                    int
		i                      int
		f                      float32
		input                  = "56.12 / 5212 / Go"
		format                 = "%f / %d / %s"
	)
	fmt.Println("Enter your full name: ")
	fmt.Scanln(&firstname, &lastname)
	fmt.Println("Enter your age: ")
	fmt.Scanln(&age)
	fmt.Printf("Hi %v %v\n", firstname, lastname)
	fmt.Sscanf(input, format, &f, &i, &s) // Sscan scans the argument string
	fmt.Println("From the string we read: ", f, i, s)
}
