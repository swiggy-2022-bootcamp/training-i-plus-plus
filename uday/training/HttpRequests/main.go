package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/hi", func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w,"Hello\n")
	})
	http.ListenAndServe(":8000",nil)
}
