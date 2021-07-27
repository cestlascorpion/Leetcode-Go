package solution

import "testing"

func TestGetMinDistance(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	target := 5
	start := 3
	exp := 1
	if GetMinDistance(nums, target, start) != exp {
		t.Fail()
	}
}
