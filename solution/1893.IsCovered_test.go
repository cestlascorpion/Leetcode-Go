package solution

import "testing"

func TestIsCovered(t *testing.T) {
	ranges := [][]int{{1, 2}, {3, 4}, {5, 6}}
	left := 2
	right := 5
	if !IsCovered(ranges, left, right) {
		t.Fail()
	}
}

func TestIsCovered2(t *testing.T) {
	ranges := [][]int{{1, 10}, {10, 20}}
	left := 21
	right := 21
	if IsCovered(ranges, left, right) {
		t.Fail()
	}
}
