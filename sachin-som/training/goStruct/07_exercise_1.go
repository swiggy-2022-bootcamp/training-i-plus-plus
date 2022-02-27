/*
vcard.go: Define a struct Address and a struct VCard. The latter contains a
person’s name, a number of addresses, a birth date, a photo. Try to find the right
data types. Make your own vcard and print its contents.
Hint: VCard must contain addresses, will they be included as values or as pointers ?
The 2nd choice is better, consuming less memory. So an Address struct with a name
and two pointers to addresses could be printed out with %v as:
{Kersschot 0x126d2b80 0x126d2be0}
*/

package main

import "fmt"

// Address Struct
type Address struct {
	address string
}

// VCard Struct
type VCard struct {
	personName      string
	personAddresses []*Address
}

func main() {
	var address1 *Address = &Address{address: "Dummy Address 1"}
	var address2 *Address = &Address{address: "Dummy Address 2"}

	var vcard VCard = VCard{}
	vcard.personName = "sachin som"
	vcard.personAddresses = append(vcard.personAddresses, address1)
	vcard.personAddresses = append(vcard.personAddresses, address2)
	fmt.Printf("%v", vcard)
}
