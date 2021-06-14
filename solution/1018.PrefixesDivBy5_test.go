package solution

import "testing"

func TestPrefixesDivBy5(t *testing.T) {
	nums := []int{0, 1, 1, 1, 1, 1}
	exp := []bool{true, false, false, false, true, false}
	res := PrefixesDivBy5(nums)
	if !IsSameBoolArray(exp, res) {
		t.Fail()
	}
}
