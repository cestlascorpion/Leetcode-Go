package solution

import "testing"

func TestSecondHighest(t *testing.T) {
	s := "dfa12321afd"
	exp := 2
	if SecondHighest(s) != exp {
		t.Fail()
	}
}

func TestSecondHighest2(t *testing.T) {
	s := "abc1111"
	exp := -1
	if SecondHighest(s) != exp {
		t.Fail()
	}
}
