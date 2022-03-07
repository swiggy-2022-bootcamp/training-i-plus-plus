package main

import (
	"fmt"
	"net/http"
	"time"
)


func main() {
	links := [] string{
		"http://google.com", 
		"http://facebook.com", 
		"http://stackoverflow.com", 
		"http://amazon.com", 
		"http://golang.org",
	}

	c := make(chan string)
	for _, link := range links {
		go checkStatus(link, c)
	}
	
	// for i := 0; i < len(links); i++ {
	// 	checkStatus(<- c)
	// }

	// //Infinite Loop
	// for {
	// 	go checkStatus(<- c, c)
	// }

	//Can also loop through chan and then run
	for l := range c{
		go func(link string) {
			//sleep the routine for 5 seconds
			time.Sleep(5 * time.Second)
			checkStatus(link, c)
		}(l)
	}
}

func checkStatus(link string, c chan string) {
	//blocking call
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " might be down")
		c <- link
	} else {
		fmt.Println(link + " is up")
		c <- link
	}
}