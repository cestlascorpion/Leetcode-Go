package solution

import "testing"

func TestFindKthPositive(t *testing.T) {
	arr := []int{2, 3, 4, 7, 11}
	k := 5
	exp := 9
	if FindKthPositive(arr, k) != exp {
		t.Fail()
	}
}
