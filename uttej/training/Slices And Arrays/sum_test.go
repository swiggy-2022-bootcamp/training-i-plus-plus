package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want int, numbers []int) {
		t.Helper()

		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	}
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 32}

		got := Sum(numbers)
		want := 42

		assertCorrectMessage(t, got, want, numbers)

	})
	// t.Run("collection of 3", func(t *testing.T) {

	// 	numbers := []int{1, 2, 3}

	// 	got := Sum(numbers)
	// 	want := 6

	// 	assertCorrectMessage(t, got, want, numbers)
	// })
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 3}, []int{4, 5, 6})
	want := []int{4, 15}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {

		t.Helper()

		if !reflect.DeepEqual(got, want) {

			t.Errorf("got %v, want %v", got, want)
		}
	}
	t.Run("summing the tails of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 4, 5}, []int{2, 4, 5})
		want := []int{9, 9}

		checkSums(t, got, want)
	})
	t.Run("summing empty slices", func(t *testing.T) {

		got := SumAllTails([]int{}, []int{2, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
