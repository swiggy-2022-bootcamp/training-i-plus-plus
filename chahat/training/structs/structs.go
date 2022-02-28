package main

import (
	"fmt"
)
type User struct {
	Name string
	Age int
}
func main(){
	fmt.Println("Structs in Go")

	chahat := User{"Chahat Bhatia",22}
	fmt.Println(chahat)
	fmt.Printf("User Details : %+v\n",chahat)
	fmt.Printf("Name is %v\n" ,chahat.Name)
}