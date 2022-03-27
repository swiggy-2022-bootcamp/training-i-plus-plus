package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader)(string,error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input),err
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Enter Name: ",reader) 
	fmt.Print(name," hello")
}