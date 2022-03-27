package main 

import "fmt"

type vertex struct{
	x int
	y int
}

func myfun(v *vertex){
	fmt.Println(*v)
}

func main(){
	
	v1:=vertex{1,2}
	fmt.Println(v1)
	myfun(&v1)
	// var pointer *vertex = &v1
	// fmt.Printf("%p\n",pointer)
	// fmt.Println(&pointer)
}