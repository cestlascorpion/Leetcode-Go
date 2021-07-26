package solution

import "testing"

func TestCheckOnesSegment(t *testing.T) {
	s := "101"
	if CheckOnesSegment(s) {
		t.Fail()
	}
}

func TestCheckOnesSegment2(t *testing.T) {
	s := "100"
	if !CheckOnesSegment(s) {
		t.Fail()
	}
}
