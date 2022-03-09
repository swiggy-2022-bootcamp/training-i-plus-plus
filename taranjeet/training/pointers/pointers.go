package main

import "fmt"

type Person struct {
	name string
	age  int
}

func changeStr1(str1 *string) {
	fmt.Printf("%s %p\n", *str1, str1)
	fmt.Println(*str1, str1)
	*str1 = "world"

}

func changeStruct(person *Person) {
	fmt.Printf("%+v %p\n", *person, person)
	fmt.Println(*person, person)
	person.age = 24
	(*person).age = 23

}

func main() {
	str1 := "hello"
	changeStr1(&str1)
	fmt.Println(str1)

	struct1 := Person{"Taranjeet", 25}
	fmt.Printf("%+v\n", struct1)

	changeStruct(&struct1)
	fmt.Printf("%+v\n", struct1)

}
