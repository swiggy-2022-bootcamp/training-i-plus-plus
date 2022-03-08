package main

import (
	"fmt"
)

type specialErrors struct {
	a, b    int
	message string
}

func (e specialErrors) Error() string {
	return fmt.Sprint("Error - ", e.message, " with args ", e.a, " ", e.b)
}

func canNotAddZero(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return -1, specialErrors{a, b, "Can not add zero"}
	} else {
		return a + b, nil
	}
}

func main() {
	res1, err := canNotAddZero(1, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res1)
	}
	res2, err := canNotAddZero(1, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res2)
	}

}
