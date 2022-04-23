package strings

import "strings"

//CompareTwoStrings ..
func CompareTwoStrings(string1 string, string2 string) int {
	return strings.Compare(string1, string2)
}

//EqualityOfStringsCaseInsensitive ..
func EqualityOfStringsCaseInsensitive(string1 string, string2 string) bool {
	return strings.EqualFold(string1, string2)
}

//EqualityOfStringsCaseSensitive ..
func EqualityOfStringsCaseSensitive(string1 string, string2 string) bool {
	return string1 == string2
}

//RepeatStringNTimes ..
func RepeatStringNTimes(str string, n int) string {
	return strings.Repeat(str, n)
}

//ContainsSubstring ..
func ContainsSubstring(string1 string, substr string) bool {
	return strings.Contains(string1, substr)
}

//ContainsAnyChar ..
func ContainsAnyChar(string1 string, chars string) bool {
	return strings.ContainsAny(string1, chars)
}

//TrimCutset ..
func TrimCutset(str string, cutset string) string {
	return strings.Trim(str, cutset)
}

//Index ..
func Index(str string, substr string) int {
	return strings.Index(str, substr)
}
