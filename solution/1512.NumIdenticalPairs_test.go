package solution

import "testing"

func TestNumIdenticalPairs(t *testing.T) {
	nums := []int{1, 2, 3, 1, 1, 3}
	if NumIdenticalPairs(nums) != 4 {
		t.Fail()
	}
}

func TestNumIdenticalPairs2(t *testing.T) {
	nums := []int{1, 2, 3, 1, 1, 3}
	if NumIdenticalPairs(nums) != 4 {
		t.Fail()
	}
}
