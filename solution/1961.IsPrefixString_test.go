package solution

import "testing"

func TestIsPrefixString(t *testing.T) {
	s := "iloveleetcode"
	words := []string{"i", "love", "leetcode"}
	if !IsPrefixString(s, words) {
		t.Fail()
	}
}

func TestIsPrefixString2(t *testing.T) {
	s := "iloveleetcod"
	words := []string{"i", "love", "leetcode"}
	if IsPrefixString(s, words) {
		t.Fail()
	}
}

func TestIsPrefixString3(t *testing.T) {
	s := "iloveleetcode"
	words := []string{"i", "love", "leetcod"}
	if IsPrefixString(s, words) {
		t.Fail()
	}
}
