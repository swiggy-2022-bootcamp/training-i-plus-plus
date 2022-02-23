package add

import "testing"

func TestAdd(t *testing.T) {
	if Add(50, 60) != 110 {
		t.Errorf("Invalid addition")
	}
}

type AddTest struct {
	a, b, expected int
}

func testSingleAdd(test AddTest, t *testing.T) {
	if res := Add(test.a, test.b); res != test.expected {
		t.Errorf("Invalid addition : expected %d actual %d", test.expected, res)
	}
}

func TestMultipleAdd(t *testing.T) {

	allTests := make([]AddTest, 3)

	allTests = append(allTests, AddTest{10, 20, 30})
	allTests = append(allTests, AddTest{20, 30, 50})
	allTests = append(allTests, AddTest{30, 50, 80})

	for _, test := range allTests {
		testSingleAdd(test, t)
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(20, 30)
	}
}
