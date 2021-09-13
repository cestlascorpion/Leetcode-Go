package solution

import "testing"

func TestMakeFancyString(t *testing.T) {
	s := "leeetcode"
	exp := "leetcode"
	if MakeFancyString(s) != exp {
		t.Fail()
	}
}

func TestMakeFancyString2(t *testing.T) {
	s := "aaabaaaa"
	exp := "aabaa"
	if MakeFancyString(s) != exp {
		t.Fail()
	}
}

func TestMakeFancyString3(t *testing.T) {
	s := "aab"
	exp := "aab"
	if MakeFancyString(s) != exp {
		t.Fail()
	}
}
