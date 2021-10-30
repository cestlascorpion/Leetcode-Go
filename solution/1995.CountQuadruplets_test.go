package solution

import "testing"

func TestCountQuadruplets(t *testing.T) {
	nums := []int{1, 1, 1, 3, 5}
	exp := 4
	if CountQuadruplets(nums) != exp {
		t.FailNow()
	}
}

func TestCountQuadruplets2(t *testing.T) {
	nums := []int{1, 1, 1, 3, 5}
	exp := 4
	if CountQuadruplets2(nums) != exp {
		t.FailNow()
	}
}

