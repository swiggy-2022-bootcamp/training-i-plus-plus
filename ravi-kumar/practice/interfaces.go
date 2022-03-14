package main

import "fmt"

//employee interface is now a contract.
//Anyone who's type is referenced with employee interface is entitled to
//implement all the functions defined in the interface
type Employee interface {
	GetName() string
}

type Engineer struct {
	name string
}

type Manager struct {
	name string
}

func (engineer *Engineer) GetName() string {
	return "Engineer: " + engineer.name
}

func (manager *Manager) GetName() string {
	return "Manager: " + manager.name
}

func printDetails(employee Employee) {
	fmt.Println(employee.GetName())
}

func main() {
	engineer := &Engineer{name: "Ravi"}
	manager := &Manager{name: "Rama"}
	printDetails(engineer)
	printDetails(manager)
}
