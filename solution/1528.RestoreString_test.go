package solution

import "testing"

func TestRestoreString(t *testing.T) {
	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	exp := "leetcode"
	res := RestoreString(s, indices)
	if len(res) != len(exp) {
		t.Fail()
	}
	for i, v := range exp {
		if res[i] != uint8(v) {
			t.Fail()
		}
	}
}
