package subtract

import (
	"reflect"
	"testing"
)

func TestSubtract(t *testing.T) {

	got := Subtract(4, 3)
	want := 1

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestAbsoluteDifference(t *testing.T) {

	got := AbsoluteDifference([]int{3, 4})
	want := 1

	if got != want {

		t.Errorf("got %d, want %d", got, want)
	}
}

func TestAbsoluteSlices(t *testing.T) {
	t.Run("Running with negative difference", func(t *testing.T) {

		got := AbsoluteSlices([]int{3, 4}, []int{5, 9})
		want := []int{1, 4}

		if !reflect.DeepEqual(got, want) {

			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("running with positive difference", func(t *testing.T) {
		got := AbsoluteSlices([]int{4, 3}, []int{9, 5})
		want := []int{1, 4}

		if !reflect.DeepEqual(got, want) {

			t.Errorf("got %v, want %v", got, want)
		}
	})

}
