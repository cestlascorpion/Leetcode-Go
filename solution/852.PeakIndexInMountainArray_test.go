package solution

import "testing"

func TestPeakIndexInMountainArray(t *testing.T) {
	arr := []int{1, 2, 3, 2}
	exp := 2
	if PeakIndexInMountainArray(arr) != exp {
		t.Fail()
	}
}
