package solution

import "testing"

func TestSumOddLengthSubArrays(t *testing.T) {
	arr := []int{10, 11, 12}
	exp := 66
	if SumOddLengthSubArrays(arr) != exp {
		t.Fail()
	}
}
