package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Temp struct {
	a int     "json:value first"
	b float64 "json:value second"
	c string  "json:value third"
}

func (t *Temp) String() string {
	return strconv.Itoa(t.a) + "\\ " + strconv.Itoa(int(t.b)) + "\\ " + t.c
}
func main() {
	t := &Temp{a: 10, b: 11.9, c: "sachin"}
	fmt.Println(t)
	rType := reflect.TypeOf(*t)
	for i := 0; i < 3; i++ {
		tags := rType.Field(i)
		fmt.Println(tags)
	}
}
