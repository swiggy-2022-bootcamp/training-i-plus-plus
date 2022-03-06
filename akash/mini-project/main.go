package main

import (
	"fmt"
	model "sample.akash.com/model"
)

func main() {
	fmt.Println("Hello ", &model.User{"Ash", "Lambert", 20})
}
