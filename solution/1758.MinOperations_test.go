package solution

import "testing"

func TestZeroOneMinOperations(t *testing.T) {
	s := "0100"
	exp := 1
	if ZeroOneMinOperations(s) != exp {
		t.Fail()
	}
}

func TestZeroOneMinOperations2(t *testing.T) {
	s := "0101"
	exp := 0
	if ZeroOneMinOperations(s) != exp {
		t.Fail()
	}
}

func TestZeroOneMinOperations3(t *testing.T) {
	s := "0000"
	exp := 2
	if ZeroOneMinOperations(s) != exp {
		t.Fail()
	}
}

