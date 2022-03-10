package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("running for a", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		if got != want {

			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("running for b", func(t *testing.T) {

		got := Repeat("b", 32)
		want := "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func ExampleRepeat() {

	got := Repeat("c", 4)

	fmt.Println(got)
	//Output: cccc
}

func BenchmarkRepeat(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}
