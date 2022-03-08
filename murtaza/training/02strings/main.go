package mystrings

import "strings"

//------- String Operation in Go
func CompareTwoStrings(str1 string, str2 string) int {
	return strings.Compare(str1, str2)
}

func EqualityOfStringsCaseInsensitive(str1 string, str2 string) bool {
	return strings.EqualFold(str1, str2)
}

func EqualityOfStringsCaseSensitive(str1 string, str2 string) bool {
	return str1 == str2
}

func RepeatStringNTimes(str string, n int) string {
	return strings.Repeat(str, n)
}

func ContainsSubstring(str1 string, substr string) bool {
	return strings.Contains(str1, substr)
}

func ContainsAnyChar(str1 string, chars string) bool {
	return strings.ContainsAny(str1, chars)
}

func TrimCutset(str string, cutset string) string {
	return strings.Trim(str, cutset)
}

func Index(str string, substr string) int {
	return strings.Index(str, substr)
}
