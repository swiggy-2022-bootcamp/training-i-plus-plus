package goTesting

import "testing"

var tests = []struct {
	text     string
	expected string
}{
	{text: "first", expected: "tsrif"},
	{text: "a", expected: "a"},
	{text: "", expected: ""},
	{text: "aaaaa", expected: "aaaaa"},
	{text: "reverse", expected: "esrever"},
}

func TestStringReverse(t *testing.T) {
	for ind, test := range tests {
		actual := ReverseString(test.text)
		if actual != test.expected {
			t.Errorf("---------------------%v test failed. Expected: (%s) but got (%s)\n", ind, test.expected, actual)
		}
	}
}
