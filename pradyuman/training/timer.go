package main

import(
	"fmt"
	"time"
)

func main(){
	
	timer1:=time.NewTimer(2*time.Second)
	<-timer1.C
	fmt.Print("timer 1 got fired\n")
}