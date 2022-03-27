package add

import(
	"testing"
	"fmt"
)

func TestAdd(t *testing.T){
	actual_res:=Add(10,20)
	expected_res:=30

	if actual_res!=expected_res{
		t.Errorf("actual is %d and expected is %d",actual_res,expected_res)
	}
}


type addTest struct{
	a,b,want int
}

func TestIntMinTableDriven(t *testing.T) {
    var tests = []addTest{
        {0, 1, 0},
        {1, 0, 0},
        {2, -2, -2},
        {0, -1, -1},
        {-1, 0, -1},
    }

    for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {
            ans := IntMin(tt.a, tt.b)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}

func BenchmarkIntMin(b *testing.B) {
		for i := 0; i < b.N; i++ {
			IntMin(1, 2)
		}
}