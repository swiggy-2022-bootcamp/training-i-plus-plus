package calculation

import "testing"

type addTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	addTest{1, 2, 3},
	addTest{2, 3, 4},
	addTest{3, 4, 5},
	addTest{4, 5, 6},
}

func TestAdd(t *testing.T) {

	for _, test := range addTests {

		if actual := Add(test.arg1, test.arg2); actual != test.expected {
			t.Errorf("Add(%d,%d) = %d; expected %d", test.arg1, test.arg2, actual, test.expected)
		}
}

func BenchmarkAdd(b *testing.B) {

	for i := 0; i < b.N; i++ {
		for _, test := range addTests {
			Add(10,10)
		}
	}

}

// func TestAdd(t *testing.T) {

// 	actual_res := Add(10, 10)
// 	expected_res := 20

// 	if actual_res != expected_res {
// 		t.Errorf("actual is %d, expected is %d", actual_res, expected_res)
// 	}
// }
