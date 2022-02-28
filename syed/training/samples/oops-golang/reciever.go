package main
import (
  "fmt"
)
// declare a custom type
type person struct {
  fname string
  lname string
}
func main() {
  // create a struct
  p := person{"ABC", "XYZ"}
  // invoke receiver function
  name := p.getName()
  fmt.Println("Full name is: " + name)
}
// receiver function
func (p person) getName(){
  return p.fname + " " + p.lname
}