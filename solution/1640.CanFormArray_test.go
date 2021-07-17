package solution

import "testing"

func TestCanFormArray(t *testing.T) {
	arr := []int{49, 18, 16}
	pieces := [][]int{{16, 18, 49}}
	if CanFormArray(arr, pieces) {
		t.Fail()
	}
}

func TestCanFormArray2(t *testing.T) {
	arr := []int{91, 4, 64, 78}
	pieces := [][]int{{78}, {4, 64}, {91}}
	if !CanFormArray(arr, pieces) {
		t.Fail()
	}
}
