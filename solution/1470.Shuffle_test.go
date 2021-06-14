package solution

import "testing"

func TestShuffle(t *testing.T) {
	nums := []int{2, 5, 1, 3, 4, 7}
	n := 3
	res := Shuffle(nums, n)
	exp := []int{2, 3, 5, 4, 1, 7}
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}
