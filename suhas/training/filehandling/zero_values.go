package main 

import "fmt"

func main() {
	var a int 
	fmt.Printf("Default zero vaue of int: ")
	fmt.Println(a)

	var b uint 
	fmt.Printf("Default zero value of uint:")
	fmt.Println(b)

	var c float32
	fmt.Printf("Default zero value of float:")
	fmt.Println(c)

	var d byte
	fmt.Printf("Default zero value of byte:")
	fmt.Println(d)

	var e bool
	fmt.Printf("Default zero value of bool:")
	fmt.Println(e)

	var f string
	fmt.Printf("Default zero value of string:")
	fmt.Println(f)

	var g complex64
	fmt.Printf("Default zero value of complex:")
	g = complex(3,4)
	fmt.Printf("value of complex :")
	fmt.Println(g)

	fmt.Println(real(g))
	fmt.Println(imag(g))


	var h [2]bool
	fmt.Println("Default zero value of an array: ")
	fmt.Println(h)

	var i [2]int
	fmt.Println("Default zero value of an array: ")
	fmt.Println(i)

	var j func()
	fmt.Println("Default zero value of an function: ")
	fmt.Println(j)

	var k []int
	fmt.Println(k == nil)
	fmt.Println("Printing slices")
	fmt.Println(k)

	type sample struct {
		a int 
		b bool
	}

	l := sample{}
	fmt.Println("Default zero value of a struct")
	fmt.Println(l)

	var m map[bool]bool
	fmt.Println(m==nil)
	fmt.Println("Priniting map:")
	fmt.Println(m)

	var n interface{}
	fmt.Println("Defaulrt value of interface")
	fmt.Println(n)

	s := '😅' //unicode
	s_rune_1 := rune(s)

	fmt.Println(s_rune_1)

	//difference b/w byte and rune
	s_1 := "GÖ"
 
    s_rune := []rune(s_1)
    s_byte := []byte(s_1)
     
    fmt.Println(s_rune)  // [71 214]
    fmt.Println(s_byte)  // [71 195 150]

	ss := "Golang"
	ss_rune := []rune(ss)
	fmt.Println(ss_rune)

	var t *int
	fmt.Println("Default value of a pointer")
	fmt.Println(t)
}