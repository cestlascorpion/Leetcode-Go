package solution

import "testing"

func TestReversePrint(t *testing.T) {
	head := constructList()
	exp := []int{5, 4, 3, 2, 1}
	res := ReversePrint(head)
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}
