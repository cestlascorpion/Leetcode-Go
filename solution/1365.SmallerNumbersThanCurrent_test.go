package solution

import "testing"

func TestSmallerNumbersThanCurrent(t *testing.T) {
	nums := []int{8, 1, 2, 2, 3}
	res := SmallerNumbersThanCurrent(nums)
	exp := []int{4, 0, 1, 1, 3}
	if len(res) != len(exp) {
		t.Fail()
	}
	for i, v := range exp {
		if res[i] != v {
			t.Fail()
		}
	}
}

func TestSmallerNumbersThanCurrent2(t *testing.T) {
	nums := []int{7, 7, 7, 7}
	res := SmallerNumbersThanCurrent(nums)
	exp := []int{0, 0, 0, 0}
	if len(res) != len(exp) {
		t.Fail()
	}
	for i, v := range exp {
		if res[i] != v {
			t.Fail()
		}
	}
}
