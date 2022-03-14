package main

import "fmt"

type Engineer struct {
	name string
}

func (engineer *Engineer) formatName(prefix string) {
	engineer.name = prefix + engineer.name
}

func (engineer *Engineer) setName(name string) {
	prefix := "Mr."
	//current value of prefix at the time of incurring deferred function's definition will be used/copied
	//i.e even if we change prefix value to something else after deferred func definition, it doesn't reflect in deferred function
	defer engineer.formatName(prefix)
	engineer.name = name
	prefix = "Master"
}

func main() {
	ravi := Engineer{name: "Ravi"}
	ravi.formatName("Mr.")
	fmt.Println("Formatted ravi: ", ravi)
	ravi.setName("RaviKumar")
	fmt.Println("Formatted RaviKumar: ", ravi)

}
