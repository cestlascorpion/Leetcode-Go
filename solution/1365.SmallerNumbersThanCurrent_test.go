package solution

import "testing"

func TestSmallerNumbersThanCurrent(t *testing.T) {
	nums := []int{8, 1, 2, 2, 3}
	res := SmallerNumbersThanCurrent(nums)
	exp := []int{4, 0, 1, 1, 3}
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}

func TestSmallerNumbersThanCurrent2(t *testing.T) {
	nums := []int{7, 7, 7, 7}
	res := SmallerNumbersThanCurrent(nums)
	exp := []int{0, 0, 0, 0}
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}
