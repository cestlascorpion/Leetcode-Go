package solution

import "testing"

func TestSubsetXORSum(t *testing.T) {
	nums := []int{5, 1, 6}
	exp := 28
	if exp != SubsetXORSum(nums) {
		t.Fail()
	}
}

func TestSubsetXORSum2(t *testing.T) {
	nums := []int{5, 1, 6}
	exp := 28
	if exp != SubsetXORSum2(nums) {
		t.Fail()
	}
}

