package solution

import "testing"

func TestAreAlmostEqual(t *testing.T) {
	s1 := "bank"
	s2 := "kanb"
	if !AreAlmostEqual(s1, s2) {
		t.Fail()
	}
}

func TestAreAlmostEqual2(t *testing.T) {
	s1 := "attack"
	s2 := "defend"
	if AreAlmostEqual(s1, s2) {
		t.Fail()
	}
}
