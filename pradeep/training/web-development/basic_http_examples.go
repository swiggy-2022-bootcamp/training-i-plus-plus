package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	port := 8080
	http.HandleFunc("/hello", helloHandler)
	log.Printf("Listening on port %v", port)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
