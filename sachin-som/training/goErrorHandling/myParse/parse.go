package myParse

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseError indicates an error in converting a word
// into an integer.
type ParseError struct {
	Index int    // Index into the space-separated list of words
	Word  string // The word that generated the parse error
	err   error
}

func (e *ParseError) String() string {
	return fmt.Sprintf("pkg parse: error parsing %q as int", e.Word)
}

// Parse parses the space-separated route words and put as integers
func Parse(input string) (numbers []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()

	fields := strings.Fields(input)
	numbers = fields2numbers(fields) // panic can occure
	return
}

func fields2numbers(fields []string) (numbers []int) {
	if len(fields) == 0 {
		panic("no words to parse")
	}

	for idx, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(&ParseError{idx, field, err})
		}
		numbers = append(numbers, num)
	}
	return
}
