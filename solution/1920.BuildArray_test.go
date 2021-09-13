package solution

import "testing"

func TestBuildAnArray(t *testing.T) {
	nums := []int{0, 2, 1, 5, 3, 4}
	exp := []int{0, 1, 2, 4, 5, 3}
	if !IsSameArray(BuildAnArray(nums), exp) {
		t.Fail()
	}
}
