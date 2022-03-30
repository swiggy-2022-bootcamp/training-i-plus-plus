package utils_test

import (
	"fmt"
	"strconv"
	"tejas/utils"
	"testing"
)

func TestGenerateRandomHash(t *testing.T) {
	var tests = []struct {
		input int
	}{
		{10},
		{20},
		{100},
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(test.input), func(t *testing.T) {
			output := utils.GenerateRandomHash(test.input)
			fmt.Println("output:", output)
			if len(output) != test.input {
				t.Errorf("GenerateRandomHash(%d) returned %d, want %d", test.input, len(output), test.input)
			}
		})
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")

func TestGenerateRandomHashQuality(t *testing.T) {
	var tests = []struct {
		input int
	}{
		{10},
		{20},
		{100},
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(test.input), func(t *testing.T) {
			output := utils.GenerateRandomHash(test.input)
			fmt.Println("output:", output)
			for _, char := range output {
				found := false
				for _, rune := range letterRunes {
					if char == rune {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("GenerateRandomHash(%d) returned %s, contains non-alphanumeric characters", test.input, output)
				}
			}
		})
	}
}
