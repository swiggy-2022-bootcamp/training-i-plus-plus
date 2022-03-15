package main

import "fmt"

type AboutAble interface {
	About()
}

type Book struct {
	name string
}

func (book *Book) About() {
	fmt.Println("Book Name: ", book.name)
}

func main() {
	//Note: When we reference a non-interface type T to an interface I, we call it dynamic typing.
	//Because, type T can be many in number for any given I

	//since Book implements all functions defined in AboutAble (i.e functions in Book are super set of functions defined in AboutAble interface)
	//it can be referenced like this.
	//A *Book value is boxed into an interface value of type AboutAble.
	var book1 AboutAble = &Book{name: "Golang"}
	fmt.Println(book1)

	//anyVal is a blank interface type. Anything can be assigned to it, as all the functions in every type we can think of is a superset of an empty function set.
	//And W.K.T a blank interface type is precisely that. A set of zero function definitions
	var anyVal interface{} = &Book{name: "Erlang"}
	fmt.Println(anyVal)
	anyVal = 12
	fmt.Println(anyVal)
	// Clear the boxed value in interface value i.
	anyVal = nil
	fmt.Println(anyVal) //prints nothing
}
