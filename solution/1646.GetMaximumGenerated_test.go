package solution

import "testing"

func TestGetMaximumGenerated(t *testing.T) {
	n := 7
	exp := 3
	if GetMaximumGenerated(n) != exp {
		t.Fail()
	}
}

func TestGetMaximumGenerated2(t *testing.T) {
	n := 2
	exp := 1
	if GetMaximumGenerated(n) != exp {
		t.Fail()
	}
}

func TestGetMaximumGenerated3(t *testing.T) {
	n := 3
	exp := 2
	if GetMaximumGenerated(n) != exp {
		t.Fail()
	}
}
