package main

import "fmt"

var userCredentials = map[string]string{}

func setUp() {
	userCredentials["John"] = "pass123"
	userCredentials["Jill"] = "321Pass"
	userCredentials["James"] = "@23Pass"
}

func authenticateUser(username string, password string) bool {
	stored_password, has_key := userCredentials[username]
	return has_key && stored_password == password
}

func deleteUser(username string) {
	delete(userCredentials, username)
}

func main() {
	setUp()
	fmt.Println(authenticateUser("John", "pass123"))
	fmt.Println(authenticateUser("John", "pass1233"))

	fmt.Println(userCredentials)
	deleteUser("Jill")
	fmt.Println("Deleted username Jill: ", userCredentials)
}
