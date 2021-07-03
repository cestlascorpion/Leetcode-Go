package solution

import "testing"

func TestCanBeEqual(t *testing.T) {
	target := []int{1, 2, 3, 4}
	arr := []int{2, 4, 1, 3}
	if !CanBeEqual(target, arr) {
		t.Fail()
	}
}


