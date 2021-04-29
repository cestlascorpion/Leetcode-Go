package solution

import "testing"

func TestMinOperations(t *testing.T) {
	nums := []int{1, 1, 1}
	exp := 3
	if MinOperations(nums) != exp {
		t.Fail()
	}
}

func TestMinOperations2(t *testing.T) {
	nums := []int{1,5,2,4,1}
	exp := 14
	if MinOperations(nums) != exp {
		t.Fail()
	}
}