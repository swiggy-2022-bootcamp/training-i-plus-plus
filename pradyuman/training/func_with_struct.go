package main 

import "fmt"

type vertex struct{
	x int
	y int
}

func main(){
	
	v1:=vertex{1,2}
	fmt.Println(v1)

	v1.coordinate()

}

func (v vertex) coordinate(){
	fmt.Printf("the co-ordinates are %v %v",v.x,v.y)
}