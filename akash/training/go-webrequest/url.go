package main

import (
	"fmt"
	"net/url"
)

const myurl2 = "http://www.lco.dev:3000/learn?coursename=python&payment=tvxnsdfBsdf2nxcv56"

func main2() {
	fmt.Println(myurl)
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	fmt.Println(result.Path)

	qparams := result.Query()
	fmt.Println(qparams)

	for key, val := range qparams {
		fmt.Println(key, val)
	}
}
