package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}
func main() {
	response, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	// fmt.Println(response.Body)

	// bs := make([]byte, 99999) //empty byte slice with 99999 space

	// //Calling read method from Reader interface
	// response.Body.Read(bs)

	// //html from google.com
	// fmt.Println(string(bs))

	lw := logWriter{}

	io.Copy(lw, response.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}