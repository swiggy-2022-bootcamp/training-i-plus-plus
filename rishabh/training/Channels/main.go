package main

import (
	"fmt"
	"net/http"
	"time"
)

var startTime time.Time

func checkWebsite(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "Site seems to be down")
		c <- link
		return
	}
	c <- link
	fmt.Println(time.Since(startTime), link, "OK")
}

func main() {
	links := []string{"https://google.com", "https://facebook.com", "https://rishabhmishra.me", "https://golang.org"}

	c := make(chan string)
	startTime = time.Now()
	for _, link := range links {
		go checkWebsite(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(time.Second)
			checkWebsite(link, c)
		}(l)
	}

}
