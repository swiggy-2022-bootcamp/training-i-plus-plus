package main

import (
	"fmt"
	"sort"
)
func main(){
	fmt.Println("Slices in Golang")

	var fruits = []string{"Apple","Grapes","Guava"}

	var veggies =[3]string{"Onion","Yam","JackFruit"}
	fmt.Printf("type %T\n",fruits)
	fmt.Printf("type %T\n",veggies)
    fmt.Println(fruits)

	fruits= append(fruits, "Mango","Banana")
   fmt.Println(fruits)

   favFruits := append(fruits[1:3])
   fmt.Println(favFruits)
   
   score := make([]int,3)
   fmt.Println(score)

   score[0]=10
   score[1]=20
   score[2]=30

   fmt.Println(score)

   score=append(score, 50,40)

   fmt.Println(score)

   sort.Ints(score)
   fmt.Println(score)
   
// remove element from slice based on index
   newslice := []int{1,23,4,3,22}
   fmt.Println(newslice)
   index:=2
   newslice = append(newslice[:index],newslice[index+1:]...)
   fmt.Println(newslice)







}
