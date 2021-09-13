package solution

import "testing"

func TestIsPrefixString(t *testing.T) {
	s := "iloveleetcode"
	words := []string{"i", "love", "leetcode"}
	if !IsPrefixString(s, words) {
		t.Failed()
	}
}

func TestIsPrefixString2(t *testing.T) {
	s := "iloveleetcod"
	words := []string{"i", "love", "leetcode"}
	if IsPrefixString(s, words) {
		t.Failed()
	}
}

func TestIsPrefixString3(t *testing.T) {
	s := "iloveleetcode"
	words := []string{"i", "love", "leetcod"}
	if IsPrefixString(s, words) {
		t.Failed()
	}
}
