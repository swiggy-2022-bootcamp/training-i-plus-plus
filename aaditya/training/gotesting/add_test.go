package calculation

import "testing"

type addTest struct{
	arg1		int
	arg2		int
	expectedRes	int
}

var addTests = []addTest{
	addTest{10,10,20},
	addTest{20,20,40},
	addTest{30,30,60},
	addTest{40,40,80},
}
func TestAdd(t *testing.T){
	actualRes:=Add(10,10)
	expectedRes := 20

	if actualRes != expectedRes {
		t.Errorf("Corect answer should be %d and found %d",expectedRes,actualRes)
	}
}

func TestAdd1(t *testing.T){
	for _,test:=range addTests{
		if result := Add(test.arg1,test.arg2);result!=test.expectedRes{
			t.Errorf("Corect answer should be %d and found %d",test.expectedRes,result)
		}
	}

}

func TestSub(t *testing.T){
	actualRes:=Sub(10,10)
	expectedRes := 0

	if actualRes != expectedRes {
		t.Errorf("Corect answer should be %d and found %d",expectedRes,actualRes)
	}
}

func BenchmarkAdd(b *testing.B){
	for i:=0; i< b.N; i++ {
		Add(10,10)
	}
}