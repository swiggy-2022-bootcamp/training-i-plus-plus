package main

import (
	"fmt"
	"strings"
)

type Student struct{ firstName, lastName string }

func toUpper(student *Student) {
	student.firstName = strings.ToUpper(student.firstName)
	student.lastName = strings.ToUpper(student.lastName)
}

func main() {

	// 1. Struct as a literal
	per1 := &Student{"sachin", "som"}

	toUpper(per1)
	fmt.Println(per1)

	// 2. Struct as a pointer
	var stu1 *Student
	stu1 = new(Student)
	stu1.firstName = "sachin"
	stu1.lastName = "som"
	toUpper(stu1)
	fmt.Println(stu1)

	// 3. Struct as a value type
	var stu2 Student
	stu2.firstName = "sachin"
	stu2.lastName = "som"
	toUpper(&stu2)
	fmt.Println(stu2)
}
