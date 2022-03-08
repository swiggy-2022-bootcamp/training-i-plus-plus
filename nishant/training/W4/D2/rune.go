package main

import (
	"fmt"
)

// Rune is alias for int32

func main() {
	s := 'z'

	fmt.Println(rune(s))

	str := "asdfasdfa"

	str_rune := []rune(str)
	fmt.Println(str_rune)

}
