package solution

import "testing"

func TestCountConsistentStrings(t *testing.T) {
	allowed := "cad"
	words := []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}
	res := CountConsistentStrings(allowed, words)
	exp := 4
	if exp != res {
		t.Fail()
	}
}

func TestCountConsistentStrings2(t *testing.T) {
	allowed := "fstqyienx"
	words := []string{"n", "eeitfns", "eqqqsfs", "i", "feniqis", "lhoa", "yqyitei", "sqtn", "kug", "z", "neqqis"}
	res := CountConsistentStrings(allowed, words)
	exp := 8
	if exp != res {
		t.Fail()
	}
}
