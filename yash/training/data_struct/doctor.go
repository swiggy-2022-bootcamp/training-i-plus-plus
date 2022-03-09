package main

import "fmt"

type doctor struct {
	name        string
	age         string
	exp         int
	rating      string
	speciality  string
	description string
	charges     int
}

func (d doctor) doctor_profile() {
	fmt.Println("The name of Doctor is ", d.name)
	fmt.Println("The age of Doctor is ", d.age)
	fmt.Println("The experience of Doctor is ", d.exp)
	fmt.Println("The speciality of Doctor is ", d.speciality)
	fmt.Println("The rating of Doctor is ", d.rating)
	fmt.Println("Description ", d.description)
	fmt.Println("Charges ", d.charges)
}
func main() {
	doctor1 := doctor{
		name:        "Dr.Rajesh",
		age:         "45",
		exp:         20,
		speciality:  "General medicines",
		rating:      "4.5",
		description: "20 years of experience in Health care. Lives in Andheri, Mumbai",
		charges:     1000,
	}
	doctor1.doctor_profile()

	
}
