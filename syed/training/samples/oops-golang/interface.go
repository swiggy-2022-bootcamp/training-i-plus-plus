package main

import (
	"fmt"
)

/*
class Animal
   virtual abstract Speak() string
*/
type Animal interface {
	Speak() string
}

/*
class Dog
  method Speak() string //non-virtual
     return "Woof!"
*/
type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

/*
class Cat
  method Speak() string //non-virtual
     return "Meow!"
*/
type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow!"
}

/*
class Llama
  method Speak() string //non-virtual
     return "LaLLamaQueLLama!"
*/
type Llama struct {
}

func (l Llama) Speak() string {
	return "LaLLamaQueLLama!"
}

/*
func main
  var animals = [ Dog{}, Cat{}, Llama{} ]
  for each animal in animals
     print animal.Speak() // method dispatch via jmp-table
*/

func main() {
	animals := []Animal{Dog{}, Cat{}, Llama{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak()) // method dispatch via jmp-table
	}
}