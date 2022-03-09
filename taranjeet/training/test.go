package main

import (
	"bufio"
	"fmt"
	"home/function"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func createMap() map[string]int {
	myMap := make(map[string]int)
	for i := 0; i < 26; i++ {
		myMap[string(rune(97+i))] = int(rune(i))
	}
	return myMap
}

func main() {
	var x, y, z int
	var a = "hello"
	b := 8.90
	x, y = 5, 12
	fmt.Println(x, y, z, a, b)

	fmt.Printf("hi %d %s %0.2f\n", 10, "wdqw", 10.1284)

	fmt.Print("Enter a number: ")

	reader, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal("Unexpected behaviour!")

	}
	fmt.Println(reflect.TypeOf(reader))

	reader = strings.TrimSpace(reader)
	readerInteger, err := strconv.ParseFloat(reader, 64)
	if err != nil {
		log.Fatal("Unexpected behaviour-2!")

	}

	fmt.Println(reflect.TypeOf(readerInteger))

	var status string

	if readerInteger > 40 {
		status = "Pass"
	} else {
		status = "Fail"
	}
	fmt.Println(reader, status)

	finalString, secString := function.TestFunction("ab", "cd", 3)
	fmt.Println(finalString, secString)

	myMap := createMap()
	value, isExist := myMap["q"]
	if isExist == false {
		fmt.Println("Element do not exist")
	} else {
		fmt.Printf("Value: %d\n", value)
	}
	fmt.Println(myMap)

}
