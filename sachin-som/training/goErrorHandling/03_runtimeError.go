// runtime errors can be created using 'panic'
package main

import "os"

var user = os.Getenv("USER")

func check() {
	if user == "" {
		panic("unknown user: user not found $USER")
	}
}

func main() {
	check()
}
