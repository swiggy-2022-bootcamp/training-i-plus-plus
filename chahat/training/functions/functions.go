package main
import (
	"fmt"
)
func Greeter(){
	fmt.Println("Hello from Golang!")
}
func Greeter1(name string){
	fmt.Printf("Welcome from Golang! %v\n",name)
}
func Greeter2(name string) string{
	return fmt.Sprintf("Hi, %s! Welcome.\n", name)
}
func main(){
	fmt.Println("Functions in Go")
	Greeter()
	Greeter1("Chahat")
	fmt.Println(Greeter2("Chahat"))
}