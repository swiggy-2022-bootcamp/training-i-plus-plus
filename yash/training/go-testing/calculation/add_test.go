package calculation

import (
	"fmt"
	"testing"
)

type Addtest struct {
	a               int
	b               int
	expected_result int
}

var addTests = []Addtest{
	{1, 2, 3},
	{2, 3, 5},
	{3, 4, 7},
}

func TestAdd(t *testing.T) {
	res := Add(10, 10)
	expected_result := 20
	if res != expected_result {
		t.Errorf("Expected result is %d, but got %d", expected_result, res)
	}

	for _, test := range addTests {
		result := Add(test.a, test.b)
		if result != test.expected_result {
			t.Errorf("Expected result is %d, but got %d", test.expected_result, result)
		}
	}
}
func BenchmarkAdd(b *testing.B) {
	// fmt.Println("Benchmarking Add()")
	fmt.Println("\nb.N is", b.N)
	for i := 0; i < b.N; i++ {

		Add(i, 100000)
	}
}
