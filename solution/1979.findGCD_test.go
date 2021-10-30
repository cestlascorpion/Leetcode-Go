package solution

import "testing"

func TestName(t *testing.T) {
	nums := []int{2, 5, 6, 9, 10}
	exp := 2
	if FindGCD(nums) != exp {
		t.FailNow()
	}
}
