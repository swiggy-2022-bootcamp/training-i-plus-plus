package main 

import(
	"fmt"
	"time"
)

func task(t chan bool){
	fmt.Print("starting...")
	time.Sleep(time.Second)
	fmt.Println("done...")

	t <- true
}

func main(){
	t := make(chan bool,1)
	go task(t)

	fmt.Print(<- t)

	go task(t)
	fmt.Print(<- t)
}