package solution

import "testing"

func TestMinTimeToType(t *testing.T) {
	word := "abc"
	exp := 5
	if MinTimeToType(word) != exp {
		t.FailNow()
	}
}

func TestMinTimeToType2(t *testing.T) {
	word := "bza"
	exp := 7
	if MinTimeToType(word) != exp {
		t.FailNow()
	}
}
