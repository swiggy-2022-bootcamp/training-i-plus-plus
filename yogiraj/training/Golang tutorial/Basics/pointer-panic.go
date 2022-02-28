package main

import "fmt"

type T struct{
	V int
	tt *T
}
func (t *T) hello() string{
	return "world"
}


func main(){
	var t *T = nil
	fmt.Println(t)
	fmt.Println(t.tt.hello())
}

