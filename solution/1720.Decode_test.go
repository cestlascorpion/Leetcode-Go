package solution

import "testing"

func TestDecode(t *testing.T) {
	encoded := []int{1, 2, 3}
	first := 1
	res := Decode(encoded, first)
	exp := []int{1, 0, 2, 1}
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}
