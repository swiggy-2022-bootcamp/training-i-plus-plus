package main

import "fmt"

type person struct {
	firstName string
	lastName string
	contact contactInfo //also be simply contactInfo where the field name is contactInfo of type contactInfo
}

type contactInfo struct {
	email string
	zipCode int
}

func main() {
	// jaithun := person{"Jaithun", "Mahira"}

	// jaithun := person{firstName: "Jaithun", lastName: "Mahira"}
	// fmt.Println(jaithun)

	// var jaithun person
	// jaithun.firstName = "Jaithun"
	// jaithun.lastName = "Mahira"
	// fmt.Println(jaithun)
	// fmt.Printf("%+v", jaithun);

	jim := person {
		firstName: "Jim",
		lastName: "Party",
		contact: contactInfo{
			email: "jim@gmail.com",
			zipCode: 12334,
		},
	}

	//pass by value, so updateName doesn't work
	// jim.updateName("Jaithun")

	// jimPointer := &jim
	// jimPointer.updateName("Jaithun")

	//The above code can also be written like below
	jim.updateName("Jaithun") //person type will be automatically converted to *person
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
func (p person) print() {
	fmt.Printf("%+v", p)
}