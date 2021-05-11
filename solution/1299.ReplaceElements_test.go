package solution

import (
	"testing"
)

func TestReplaceElements(t *testing.T) {
	nums := []int{17, 18, 5, 4, 6, 1}
	res := ReplaceElements(nums)
	exp := []int{18, 6, 6, 6, 1, -1}
	if len(res) != len(exp) {
		t.Fail()
	}
	for i, v := range exp {
		if res[i] != v {
			t.Fail()
		}
	}
}
