package main

import (
	"fmt"
	"os"
)

func main(){
	// panic("There is a problem")

	_, err := os.Create("/tmp/file")
	
	if err != nil{
		fmt.Println("Some problem occured")
		panic(err)
	}
}