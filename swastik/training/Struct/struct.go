package main

import (
    "fmt"
)

type Student struct{
	name string
	age int
	class int
	rollNumber int
}

//receiver function
func (s Student) print(){
	fmt.Println("Name: ", s.name)
	fmt.Println("Age: ", s.age)
	fmt.Println("Class: ", s.class)
	fmt.Println("roll number: ", s.rollNumber)
}

func main() {
	var student1 = Student{"Rahul", 18, 12, 23}
	fmt.Println(student1)
	student1.print()
}
