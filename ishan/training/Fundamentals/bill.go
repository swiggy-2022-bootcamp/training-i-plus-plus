package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   0,
	}
	return b
}

func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v .....$%v \n", k+":", v)
		total += v
	}
	fs += fmt.Sprintf("%-25v .....$%0.2f", "total:", total)
	return fs
}

func createBill() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter bill name : ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	b := newBill(name)
	fmt.Printf("New Bill Created with name : %v \n", b.name)
}
