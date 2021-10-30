package solution

import "testing"

func TestCountKDifference(t *testing.T) {
	nums := []int{3, 2, 1, 5, 4}
	k := 2
	exp := 3
	if CountKDifference(nums, k) != exp {
		t.FailNow()
	}
}
