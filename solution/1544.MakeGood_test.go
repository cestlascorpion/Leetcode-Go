package solution

import "testing"

func TestMakeGood(t *testing.T) {
	s := "leEeetcode"
	exp := "leetcode"
	if MakeGood(s) != exp {
		t.Fail()
	}
}

func TestMakeGood2(t *testing.T) {
	s := "abBAcC"
	exp := ""
	if MakeGood(s) != exp {
		t.Fail()
	}
}
