package solution

import "testing"

func TestNumberOfMatches(t *testing.T) {
	n := 14
	exp := 13
	if NumberOfMatches(n) != exp {
		t.Fail()
	}
}

func TestNumberOfMatches2(t *testing.T) {
	n := 7
	exp := 6
	if NumberOfMatches(n) != exp {
		t.Fail()
	}
}
