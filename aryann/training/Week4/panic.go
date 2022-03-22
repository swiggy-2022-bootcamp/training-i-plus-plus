package main

import "os"

func main() {

	panic("a problem")

	err := os.Create("/tmp/file")

	if err != nil {
		panic(err)
	}
}
