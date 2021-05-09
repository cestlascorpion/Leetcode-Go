package solution

import "testing"

func TestSumBase(t *testing.T) {
	n := 34
	k := 6
	exp := 9
	if SumBase(n, k) != exp {
		t.Fail()
	}
}

func TestSumBase2(t *testing.T) {
	n := 10
	k := 10
	exp := 1
	if SumBase(n, k) != exp {
		t.Fail()
	}
}
