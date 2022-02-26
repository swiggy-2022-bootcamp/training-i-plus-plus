package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("In main.")
	go shortWait()
	go longWait()
	fmt.Println("About to sleep in main for 10s.")
	time.Sleep(10 * 1e9) // Sleep for 10 s // if we don't wait for other goRoutines program exist and goRoutine will die
	fmt.Println("Ending main.")
}

func longWait() {
	fmt.Println("Beginining long wait.")
	time.Sleep(5 * 1e9) // Sleep for 5 s
	fmt.Println("Ending long wait.")
}

func shortWait() {
	fmt.Println("Beginining short wait.")
	time.Sleep(2 * 1e9) // Sleep for 5 s
	fmt.Println("Ending short wait.")
}
