package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string
}

type helloWorldComplexResponse struct {
	// to change the name of the field to smallcase letters.
	Message string `json:"message"`
	// smallcase letters fields will not be exported by Marshal.
	Author string
}

func helloWorldComplexResponseHandler(w http.ResponseWriter, r *http.Request) {
	complexResponse := helloWorldComplexResponse{Message: "Hello World", Author: "Pradeep"}
	data, err := json.Marshal(complexResponse) // data is a byte array.
	if err != nil {
		panic("Oops")
	}
	fmt.Fprintf(w, string(data))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello World"}

	data, err := json.Marshal(response)

	if err != nil {
		panic("Oops")
	}
	fmt.Fprintf(w, string(data))
}

func main() {
	port := 8080
	http.HandleFunc("/hello", helloWorldHandler)
	http.HandleFunc("/hellocomplex", helloWorldComplexResponseHandler)
	log.Printf("Listening on port %v", port)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
