package solution

import "testing"

func TestMergeAlternately(t *testing.T) {
	word1 := "abc"
	word2 := "pqr"
	exp := "apbqcr"
	res := MergeAlternately(word1, word2)
	if !IsSameString(exp, res) {
		t.Fail()
	}
}

func TestMergeAlternately2(t *testing.T) {
	word1 := "ab"
	word2 := "pqrs"
	exp := "apbqrs"
	res := MergeAlternately(word1, word2)
	if !IsSameString(exp, res) {
		t.Fail()
	}
}
