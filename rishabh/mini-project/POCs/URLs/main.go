package main

import (
	"fmt"
	"net/url"
)

func main() {
	const giturl = "https://github.com/Rishabh3321/?tab=repositories"
	fmt.Println(giturl)

	result, _ := url.Parse(giturl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Query())
}
