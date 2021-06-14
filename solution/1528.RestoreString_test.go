package solution

import "testing"

func TestRestoreString(t *testing.T) {
	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	exp := "leetcode"
	res := RestoreString(s, indices)
	if !IsSameString(exp, res) {
		t.Fail()
	}
}
