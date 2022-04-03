package math_test

import (
	"sonarcubeApp/math"
	"testing"
)

func TestAdd(t *testing.T) {
	var res int
	res = math.Add(1, 2)
	if res != 3 {
		t.Errorf("Add(1, 2) = %d, want 3", res)
	}
}

func TestSub(t *testing.T) {
	var res int
	res = math.Sub(1, 2)
	if res != -1 {
		t.Errorf("Sub(1, 2) = %d, want -1", res)
	}
}

func TestMul(t *testing.T) {
	var res int
	res = math.Mul(1, 2)
	if res != 2 {
		t.Errorf("Mul(1, 2) = %d, want 2", res)
	}
}

func TestDiv(t *testing.T) {
	var res int
	res = math.Div(1, 2)
	if res != 0 {
		t.Errorf("Div(1, 2) = %d, want 0", res)
	}
}
