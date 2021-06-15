package solution

import "testing"

func TestFindSpecialInteger(t *testing.T) {
	arr := []int{1, 2, 2, 6, 6, 6, 6, 7, 10}
	exp := 6
	if FindSpecialInteger(arr) != exp {
		t.Fail()
	}
}
