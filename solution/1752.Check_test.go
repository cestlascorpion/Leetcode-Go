package solution

import "testing"

func TestCheck(t *testing.T) {
	nums := []int{3, 4, 5, 1, 2}
	if !Check(nums) {
		t.Fail()
	}
}
