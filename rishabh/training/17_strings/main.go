package main

import (
	"fmt"
	"unicode/utf8"
)

func runeForLoopsExample() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	for _, ch := range sample {
		fmt.Print(ch, " ")
	}
	fmt.Println("")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Println("")

}

func main() {
	runeForLoopsExample()

	const s = "สวัสดี"
	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
