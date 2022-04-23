package main

import "fmt"

//aa ..
type aa interface {
	About()
}

//Book ..
type Book struct {
	name string
}

//About ..
func (book *Book) About() {
	fmt.Println("Book Name: ", book.name)
}

func main() {

	var book1 aa = &Book{name: "Golang"}
	fmt.Println(book1)
	var anyVal interface{} = &Book{name: "Erlang"}
	fmt.Println(anyVal)
	anyVal = 12
	fmt.Println(anyVal)
	anyVal = nil
	fmt.Println(anyVal)
}
