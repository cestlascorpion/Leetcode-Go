package solution

import "testing"

func TestKLengthApart(t *testing.T) {
	nums := []int{1, 0, 0, 0, 1, 0, 0, 1}
	k := 2
	if !KLengthApart(nums, k) {
		t.Fail()
	}
}

func TestKLengthApart2(t *testing.T) {
	nums := []int{1, 0, 0, 1, 0, 1}
	k := 2
	if KLengthApart(nums, k) {
		t.Fail()
	}
}
