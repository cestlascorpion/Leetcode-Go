package solution

import "testing"

func TestCountOdds(t *testing.T) {
	low := 3
	high := 7
	exp := 3
	if CountOdds(low, high) != exp {
		t.Fail()
	}
}
