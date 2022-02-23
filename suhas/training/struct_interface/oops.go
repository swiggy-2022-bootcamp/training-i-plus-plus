package main 

import "fmt"


type animal struct {
	name string	
	characterstics []string
}

type plant struct {
	name string
	characterstics []string
}



func main() {
	animal1 := animal {
		name: "Elephant",
	}
	plant1 := plant {
		name: "paddy",
	}

	animal1.run()
	plant1.run()
	run()
}

//function refernce to a structure
func (a animal) run() {
	fmt.Println(a.name,a.characterstics," is a lazy animal, so cannot run!!")
}

func (p plant) run() {
	fmt.Println(p.name,"is a plant , hence cannot run!!")
}

func run() {
	fmt.Printf("!!\n")
}