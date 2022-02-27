package main

import "fmt"

func main(){
	var nameAgeMap map[string]int
	nameAgeMap = map[string]int{
		"James": 50,
		"ali": 32,
	}

	fmt.Println("Age of James: ", nameAgeMap["James"])

	for key, value := range nameAgeMap {
		fmt.Printf("%v is %d years old\n", key, value)
	}

	nameAgeMap["Yogi"] = 21;
	for key, value := range nameAgeMap {
		fmt.Printf("%v is %d years old\n", key, value)
	}
}