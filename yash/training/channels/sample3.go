package main

import "fmt"

func task(temp string) {
	fmt.Println("Task is running ", temp)
	// time.Sleep(time.Second * 2)
	fmt.Println("Task is done ", temp)
	// t <- true
}

func main() {
	// t := make(chan bool, 1)
	go task("0")
	for i := 0; i < 10; i++ {
		temp := fmt.Sprintf("%d", i)
		fmt.Println(i)
		go task(temp)
	}

}
