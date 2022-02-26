package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const myurl = "https://catfact.ninja/fact"

func main() {

	reponse, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", reponse)
	defer reponse.Body.Close()

	data, _ := ioutil.ReadAll(reponse.Body)
	fmt.Println(string(data))
}
