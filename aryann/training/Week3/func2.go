package main

import "fmt"

func passenger(fname string, lname string) string {

	return fname + " " + lname
}

func train(src string, dst string) string {

	return "from: " + src + " to " + dst
}

func main() {

	fmt.Println("Enter Your First Name: ")
	var fname string
	fmt.Scanln(&fname)

	fmt.Println("Enter Your Last Name: ")
	var lname string
	fmt.Scanln(&lname)

	fmt.Println("Passenger name: ", passenger(fname, lname))

	var source string = "Mumbai"
	var destination string = "Bangalore"

	fmt.Println("Train: ", train(source, destination))
}
