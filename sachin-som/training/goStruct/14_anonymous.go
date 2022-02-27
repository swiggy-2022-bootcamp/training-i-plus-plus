/*
Sometimes it can be useful to have structs which contain one or more anonymous (or embedded) fields,
that is fields with no explicit name. Only the type of such a field is mandatory and the type is then
also its name. Such an anonymous field can also be itself a struct: structs can contain embedded structs.
*/

package main

import "fmt"

type inner struct {
	field1 int
	field2 int
}

type outer struct {
	field3 int
	field4 float32
	inner  // Anonymous field name. // Duplicate Anonymous field not allowed.
}

func main() {
	o := outer{field3: 3, field4: 3.4, inner: inner{field1: 10, field2: 20}}
	fmt.Println(o)
}
