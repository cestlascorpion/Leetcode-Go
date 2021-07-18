package solution

import "testing"

func TestTotalMoney(t *testing.T) {
	n := 4
	exp := 10
	if TotalMoney(n) != exp {
		t.Fail()
	}
}

func TestTotalMoney2(t *testing.T) {
	n := 10
	exp := 37
	if TotalMoney(n) != exp {
		t.Fail()
	}
}
