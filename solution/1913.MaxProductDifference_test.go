package solution

import "testing"

func TestMaxProductDifference(t *testing.T) {
	nums := []int{4, 2, 5, 9, 7, 4, 8}
	exp := 64
	if MaxProductDifference(nums) != exp {
		t.Fail()
	}
}
