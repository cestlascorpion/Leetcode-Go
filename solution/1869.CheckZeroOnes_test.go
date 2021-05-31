package solution

import "testing"

func TestCheckZeroOnes(t *testing.T) {
	s := "1101"
	exp := true
	if CheckZeroOnes(s) != exp {
		t.Fail()
	}
}

func TestCheckZeroOnes2(t *testing.T) {
	s := "111000"
	exp := false
	if CheckZeroOnes(s) != exp {
		t.Fail()
	}
}

func TestCheckZeroOnes3(t *testing.T) {
	s := "110100010"
	exp := false
	if CheckZeroOnes(s) != exp {
		t.Fail()
	}
}

func TestCheckZeroOnes4(t *testing.T) {
	s := "1"
	exp := true
	if CheckZeroOnes(s) != exp {
		t.Fail()
	}
}
