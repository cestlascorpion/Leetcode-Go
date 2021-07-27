package solution

import "testing"

func TestCountGoodSubstrings(t *testing.T) {
	s := "xyzzaz"
	exp := 1
	if CountGoodSubstrings(s) != exp {
		t.Fail()
	}
}

func TestCountGoodSubstrings2(t *testing.T) {
	s := "aababcabc"
	exp := 4
	if CountGoodSubstrings(s) != exp {
		t.Fail()
	}
}
