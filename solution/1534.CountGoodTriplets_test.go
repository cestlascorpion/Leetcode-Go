package solution

import "testing"

func TestCountGoodTriplets(t *testing.T) {
	arr := []int{3, 0, 1, 1, 9, 7}
	a, b, c := 7, 2, 3
	exp := 4
	if CountGoodTriplets(arr, a, b, c) != exp {
		t.Fail()
	}
}
