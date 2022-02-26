package main

import (
	"fmt"
)
func main(){
	fmt.Println("Maps in Go")
	languages := make(map[string]int)
	languages["Java"]=1
	languages["Ruby"]=2
	languages["Python"]=3
	fmt.Print(languages)
	fmt.Println("\n",languages["Ruby"])

	// delete

	delete(languages,"Ruby")
	fmt.Println(languages)

	//loop
  
	for key,value := range languages {
		fmt.Println(key," : ",value)
	}

	for key,_ :=range languages {
		fmt.Println(key)
	}







}