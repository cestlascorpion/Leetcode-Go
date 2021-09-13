package solution

import "testing"

func TestAreOccurrencesEqual(t *testing.T) {
	s := "abacbc"
	if !AreOccurrencesEqual(s) {
		t.Fail()
	}
}

func TestAreOccurrencesEqual2(t *testing.T) {
	s := "aaabb"
	if AreOccurrencesEqual(s) {
		t.Fail()
	}
}
