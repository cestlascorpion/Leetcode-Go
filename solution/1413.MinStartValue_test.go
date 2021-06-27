package solution

import "testing"

func TestMinStartValue(t *testing.T) {
	nums := []int{-3, 2, -3, 4, 2}
	exp := 5
	if MinStartValue(nums) != exp {
		t.Fail()
	}
}
