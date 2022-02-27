package main

import "fmt"

type Person struct {
	name string
	age  int
	bill int
}

type Account struct {
	members []Person
	owner   Person
}

type Billable interface {
	getBill() int
}

func (p Person) getBill() int {
	return p.bill
}

func (a Account) getBill() (total_bill int) {
	for _, member := range a.members {
		total_bill += member.bill
	}
	return
}

func createAccount(p Person) Account {
	persons := []Person{p}
	return Account{
		persons,
		p,
	}
}

func (acc *Account) addMember(p Person) {
	acc.members = append(acc.members, p)
}

func main() {
	A := Person{"A", 30, 100}
	accA := createAccount(A)
	// fmt.Printf("Owner = %v, members = %v  \n\n", accA.owner, accA.members)

	B := Person{"B", 20, 200}
	accA.addMember(B)
	// fmt.Printf("Owner = %v, members = %v  \n\n", accA.owner, accA.members)

	C := Person{"C", 15, 50}
	accA.addMember(C)

	fmt.Printf("Owner = %v, members = %v  \n\n", accA.owner, accA.members)

	fmt.Printf("Bill for person : %s = %d\n", A.name, A.getBill())
	fmt.Printf("Bill for person : %s = %d\n", B.name, B.getBill())
	fmt.Printf("Bill for person : %s = %d\n", C.name, C.getBill())
	fmt.Printf("tTotal Bill for Account %d \n\n", accA.getBill())

}
