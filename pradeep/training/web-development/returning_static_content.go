package main

import (
	"net/http"
	"log"
)

var imgHandler = http.FileServer(http.Dir("./images"))
func main(){
	port := 8080
	http.Handle("/cat/", http.StripPrefix("/cat/", imgHandler))
	log.Printf("Listening on port %v", port)
	log.Fatal(http.ListenAndServe(":8080", nil))
}