package solution

import "testing"

func TestMaximumUnits(t *testing.T) {
	boxTypes := [][]int{{1, 3}, {2, 2}, {3, 1}}
	truckSize := 4
	exp := 8
	if MaximumUnits(boxTypes, truckSize) != exp {
		t.Fail()
	}
}

func TestMaximumUnits2(t *testing.T) {
	boxTypes := [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}
	truckSize := 10
	exp := 91
	if MaximumUnits(boxTypes, truckSize) != exp {
		t.Fail()
	}
}
