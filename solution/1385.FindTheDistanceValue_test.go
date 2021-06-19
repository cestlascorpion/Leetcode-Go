package solution

import "testing"

func TestFindTheDistanceValue(t *testing.T) {
	arr1 := []int{4, 5, 6}
	arr2 := []int{10, 9, 1, 8}
	d := 2
	exp := 2
	if FindTheDistanceValue(arr1, arr2, d) != exp {
		t.Fail()
	}
}
