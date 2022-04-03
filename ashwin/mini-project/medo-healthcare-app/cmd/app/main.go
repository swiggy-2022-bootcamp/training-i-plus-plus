package main

import (
	"fmt"
	"log"
	"medo-healthcare-app/cmd/router"
	"net/http"
)

func main() {
	fmt.Println("Hey, there !")
	fmt.Println("Medo - The Onestop Healthcare Point !")
	route := router.Router()
	log.Fatal(http.ListenAndServe(":9001", route))
}
