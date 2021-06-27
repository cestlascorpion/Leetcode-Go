package solution

import "testing"

func TestMinSubsequence(t *testing.T) {
	nums := []int{4, 3, 10, 9, 8}
	exp := []int{10, 9}
	res := MinSubsequence(nums)
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}
