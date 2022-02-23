package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Account struct {
	members []Person
	owner   Person
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
	A := Person{"A", 30}
	accA := createAccount(A)
	fmt.Printf("Owner = %v, members = %v  \n\n", accA.owner, accA.members)

	B := Person{"B", 20}
	accA.addMember(B)
	fmt.Printf("Owner = %v, members = %v  \n\n", accA.owner, accA.members)

	C := Person{"C", 15}
	accA.addMember(C)
	fmt.Printf("Owner = %v, members = %v  \n\n", accA.owner, accA.members)
}
