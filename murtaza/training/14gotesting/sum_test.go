package main

import (
	"fmt"
	"testing"
)

type TestObj struct {
	a        int
	b        int
	expected int
}

func TestSumPositiveFlow(t *testing.T) {
	testObj := TestObj{1, 2, 3}
	if Add(testObj.a, testObj.b) != testObj.expected {
		t.Fail()
	}
}

func TestSumNegativeFlow(t *testing.T) {
	testObj := TestObj{1, 2, 5}
	if Add(testObj.a, testObj.b) == testObj.expected {
		t.Fail()
	}
}

func BenchmarkF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(Add(1, 2))
	}
}
