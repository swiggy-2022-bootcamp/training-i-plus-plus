package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://ipp.swiggy.in"

func main() {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	} else {
		fmt.Println(response)
	}

	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	} else {
		content := string(data)
		fmt.Println(content)
	}
}
