package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	msgs := make(chan string)

	var wg sync.WaitGroup

	go func() {

		msgs <- "hello"
	}


}