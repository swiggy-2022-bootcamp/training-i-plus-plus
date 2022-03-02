package main

import "fmt"

func sequencer(start int,end int) func() int {
	i:= start-1
	return func () int {
		if i<end{
			i++
		}
		return i
	}
}

func main(){
	times:= 5
	seq:= sequencer(0,times-2)

	for times>0 {
		fmt.Println(seq())
		times--;
	}
}