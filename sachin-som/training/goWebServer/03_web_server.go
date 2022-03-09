package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	fmt.Println("Inside HEllo Server handler")
	io.WriteString(w, req.URL.Path[1:])
}

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
			os.Exit(-1)
		}
	}()
	// http.HandleFunc("/", HelloServer)
	http.Handle("/", http.HandlerFunc(HelloServer))
	err := http.ListenAndServe("localhost:5000", nil)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
