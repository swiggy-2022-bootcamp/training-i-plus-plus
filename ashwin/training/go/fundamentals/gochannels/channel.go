package main

import (
	"fmt"
	"net/http"
	"time"
)

var startTime time.Time

//WebCheck ..
func WebCheck(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "Uhho ! We are down ! :(")
		c <- link
		return
	}
	c <- link
	fmt.Println(time.Since(startTime), link, "Alright !!!")
}

func main() {
	links := []string{"https://google.com", "https://twitter.com", "https://go.dev"}

	c := make(chan string)
	startTime = time.Now()
	for _, link := range links {
		go WebCheck(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(time.Second)
			WebCheck(link, c)
		}(l)
	}

}
