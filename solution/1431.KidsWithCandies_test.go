package solution

import "testing"

func TestKidsWithCandies(t *testing.T) {
	candies := []int{2, 3, 5, 1, 3}
	extra := 3
	res := KidsWithCandies(candies, extra)
	exp := []bool{true, true, true, false, true}
	if !IsSameBoolArray(exp, res) {
		t.Fail()
	}
}
