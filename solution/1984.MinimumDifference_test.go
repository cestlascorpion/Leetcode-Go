package solution

import "testing"

func TestMinimumDifference(t *testing.T) {
	nums := []int{9, 4, 1, 7}
	k := 2
	exp := 2
	if MinimumDifference(nums, k) != exp {
		t.FailNow()
	}
}
