package main

import "fmt"

//context: if a type T implements every function defined in interface type I, it means T can be assigned to I - as they are compatible
//Since any type you take, always implements every function defined in a blank interface value (a interface with no functions defined/declared),
//we can say that they all can be assigned to/ referenced by the blank interface type.

//Congrats, you've just discovered a new type called "any" in many other languages!!!
//the reason println can print any value is because it takes a blank interface as it's parameter

type Engineer struct {
	name string
}

//blank interface value/type - a simulated custom "any" type
type Any interface {
}

//takes any type as an arguement
func doSomethingCoooool(any Any) {
	fmt.Println(any)
}

func main() {
	engineer := &Engineer{name: "Ravi"}
	doSomethingCoooool(engineer)
	doSomethingCoooool(1)
	doSomethingCoooool("mama")
}
