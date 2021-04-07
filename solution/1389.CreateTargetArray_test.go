package solution

import (
	"testing"
)

func TestCreateTargetArray(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4}
	index := []int{0, 1, 2, 2, 1}
	res := CreateTargetArray(nums, index)
	exp := []int{0, 4, 1, 3, 2}
	if len(res) != len(exp) {
		t.Fail()
	}
	for i, v := range exp {
		if res[i] != v {
			t.Fail()
		}
	}
}
