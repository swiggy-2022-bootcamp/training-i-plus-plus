package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type scrapper struct {
	filename string
}

func (s scrapper) Write(p []byte) (int, error) {
	file, _ := os.Create(s.filename)
	return file.Write(p)
}

func main() {

	resp, err := http.Get("https://rishabhmishra.me")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	var s = scrapper{filename: "doc.html"}
	io.Copy(s, resp.Body)
}
