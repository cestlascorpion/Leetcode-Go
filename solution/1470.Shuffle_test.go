package solution

import "testing"

func TestShuffle(t *testing.T) {
	nums := []int{2, 5, 1, 3, 4, 7}
	n := 3
	res := Shuffle(nums, n)
	exp := []int{2, 3, 5, 4, 1, 7}
	if len(res) != len(exp) {
		t.Fail()
	}
	for i, v := range exp {
		if res[i] != v {
			t.Fail()
		}
	}
}
