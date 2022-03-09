package calculation

import "testing"

type addTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	{10, 10, 20},
	{20, 10, 30},
	{30, 10, 40},
	{40, 10, 50},
	{50, 10, 60},
}

func TestAdd(t *testing.T) {

	for _, test := range addTests {
		if result := Add(test.arg1, test.arg2); result != test.expected {
			t.Errorf("Expected: %v, Actual: %v", test.expected, result)
		}
	}
}

func BenchmarkAdd(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Add(10, 10)
	}
}

// func TestAdd(t *testing.T) {

// 	actualRes := Add(10, 10)
// 	expectedRes := 20

// 	if actualRes != expectedRes {
// 		t.Errorf("Expected: %v, Actual: %v", actualRes, expectedRes)
// 	}
// }

// func TestSubtract(t *testing.T) {

// 	actualRes := Subtract(10, 10)
// 	expectedRes := 0

// 	if actualRes != expectedRes {
// 		t.Errorf("Expected: %v, Actual: %v", actualRes, expectedRes)
// 	}
// }
