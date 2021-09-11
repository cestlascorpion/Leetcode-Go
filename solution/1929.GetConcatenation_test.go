package solution

import "testing"

func TestGetConcatenation(t *testing.T) {
	nums := []int{1, 2, 1}
	exp := []int{1, 2, 1, 1, 2, 1}
	if IsSameArray(GetConcatenation(nums), exp) {
		t.Failed()
	}
}
