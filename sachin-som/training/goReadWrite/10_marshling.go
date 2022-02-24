// data structure -> specific format : marshelling
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Address struct {
	Type    string "json:Type of address"
	City    string "json:City"
	Country string "json:Country"
}

type Vcard struct {
	Firstname string
	Lastname  string
	addresses []*Address
	Remark    string
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	privateA := &Address{"private", "jaipur", "india"}
	workA := &Address{"work", "bangalore", "india"}
	vcard := &Vcard{"sachin", "som", []*Address{privateA, workA}, "none"}
	fmt.Println(*vcard) //{sachin som [0xc00007e480 0xc00007e4b0] none}

	// JSON: Format
	jsonVcard, _ := json.Marshal(vcard)
	fmt.Println(jsonVcard) // bytes format

	// ENCODING
	file, err := os.OpenFile("vard.json", os.O_CREATE|os.O_WRONLY, 0)
	check(err)
	defer file.Close()

	enc := json.NewEncoder(file) // takes an io.Writer implemented stream
	check(enc.Encode(vcard))     // writes the json data to the io.Writer implemented stream // takes emtpy interface

	// Summary of the above example
	// struct (data structure) ==> JSON (encoding format)
}
