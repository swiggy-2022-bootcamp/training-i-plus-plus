package main

import (
	"fmt"
	"reflect"
)

type TageType struct {
	field1 string "An Important Name"
	field2 int    "Quantity of an thing"
}

func main() {
	tt := &TageType{"Sachin Som", 4}
	structV := reflect.ValueOf(*tt)
	for i := 0; i < structV.NumField(); i++ {
		reftage(tt, i)
	}
}

func reftage(tt *TageType, ind int) {
	ttType := reflect.TypeOf(*tt)
	f := ttType.Field(ind)
	fmt.Println(f.Tag)
}
