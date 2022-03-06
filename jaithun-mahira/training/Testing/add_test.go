package calculation

import "testing"

type addTest struct{
	arg1,arg2, expected int
}

var addTests = []addTest{
	addTest{10,10,20},
	addTest{20,20,40},
	addTest{30,30,60},
	addTest{40,40,80},
	addTest{50,50,100},

}
func TestAdd(t *testing.T){
	for _, test := range addTests{
		if result := Add(test.arg1,test.arg2);result != test.expected{
			t.Errorf("actual is %d,expected is %d",result,test.expected)
			}
	}
}

func BenchmarkAdd(b *testing.B){
	for i := 0; i < b.N; i++{
		Add(10,10)
	}
}
// syntax: Benchmarkxxx(*testing.B)
// func TestAdd(t *testing.T){
// 	actual_res := Add(10,10)
// 	expected_res := 20

// 	if actual_res != expected_res{
// 		t.Errorf("actual is %d,expected is %d",actual_res,expected_res)
// 	}
// }


func TestSub(t *testing.T){
	actual_res := Sub(30,10)
	expected_res := 20

	if actual_res != expected_res{
		t.Errorf("actual is %d,expected is %d",actual_res,expected_res)
	}
}

