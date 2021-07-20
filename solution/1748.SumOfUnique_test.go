package solution

import "testing"

func TestSumOfUnique(t *testing.T) {
	nums := []int{1, 2, 3, 2}
	exp := 4
	if SumOfUnique(nums) != exp {
		t.Fail()
	}
}
