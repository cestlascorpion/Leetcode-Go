package solution

import "testing"

func TestMaxScore(t *testing.T) {
	s := "011101"
	exp := 5
	if MaxScore(s) != exp {
		t.Fail()
	}
}
