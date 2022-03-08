package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	serverPath = "http://localhost:5000/sachin"
)

func main() {
	res, err := http.Get(serverPath)
	checkError(err)
	data, err := ioutil.ReadAll(res.Body)
	checkError(err)
	fmt.Println(data)
}

func checkError(err error) {
	log.Fatal(err.Error())
}
