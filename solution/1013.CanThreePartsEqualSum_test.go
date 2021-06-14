package solution

import "testing"

func TestCanThreePartsEqualSum(t *testing.T) {
	arr := []int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}
	if !CanThreePartsEqualSum(arr) {
		t.Fail()
	}
}

func TestCanThreePartsEqualSum2(t *testing.T) {
	arr := []int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1}
	if CanThreePartsEqualSum(arr) {
		t.Fail()
	}
}

func TestCanThreePartsEqualSum3(t *testing.T) {
	arr := []int{1, -1, 1, -1}
	if CanThreePartsEqualSum(arr) {
		t.Fail()
	}
}
