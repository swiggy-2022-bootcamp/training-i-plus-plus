package main

import (
	"fmt"
)

func sayGreetings(n string){
	fmt.Printf("Hello %v\n",n);
}

func cycleName(n []string,f func(string)){
	for _, v := range n{
		f(v)
	}
}

func add(a int,b int)int{
	return a+b
}

func main() {
	fmt.Println("Hello World!")
	sayGreetings("vishu")
	cycleName([]string{"Pradyuman","swapnil"},sayGreetings)
	fmt.Println(add(1,2))
}