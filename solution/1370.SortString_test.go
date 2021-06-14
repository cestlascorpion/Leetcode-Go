package solution

import (
	"testing"
)

func TestSortString(t *testing.T) {
	s := "aaaabbbbcccc"
	res := SortString(s)
	exp := "abccbaabccba"
	if !IsSameString(exp, res) {
		t.Fail()
	}
}

func TestSortString2(t *testing.T) {
	s := "rat"
	res := SortString(s)
	exp := "art"
	if !IsSameString(exp, res) {
		t.Fail()
	}
}

func TestSortString3(t *testing.T) {
	s := "leetcode"
	res := SortString(s)
	exp := "cdelotee"
	if !IsSameString(exp, res) {
		t.Fail()
	}
}

func TestSortString4(t *testing.T) {
	s := "ggggggg"
	res := SortString(s)
	exp := "ggggggg"
	if !IsSameString(exp, res) {
		t.Fail()
	}
}
