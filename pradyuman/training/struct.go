package main 

import "fmt"

type animal struct{
	name string
	charactersitics []string
}

type vertex struct{
	x int
	y int
}

func main(){
	animal1:= animal{
		name: "bruno",
		charactersitics:[]string{"puppy","sleeps alot",},
	}
	for index,v :=  range animal1.charactersitics{
		fmt.Println(index,v)
	}

	v1:=vertex{1,2}
	fmt.Println(v1)
}