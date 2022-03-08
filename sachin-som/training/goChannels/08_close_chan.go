package main

import "fmt"

func sendData(ch chan<- string) {
	ch <- "DKJFDKf"
	ch <- "SJFIEUF"
	ch <- "WIDNFKE"
	ch <- "IEHFNEE"
	close(ch) // send signal for no data
}

func getData(ch <-chan string) {
	for {
		d, ok := <-ch // check if data received or not
		if !ok {
			break
		}
		fmt.Println(d)

	}
}
func main() {
	ch := make(chan string)
	go sendData(ch)
	getData(ch)
}
