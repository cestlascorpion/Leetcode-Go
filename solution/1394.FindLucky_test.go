package solution

import "testing"

func TestFindLucky(t *testing.T) {
	arr := []int{2, 2, 3, 4}
	exp := 2
	if FindLucky(arr) != exp {
		t.Fail()
	}
}
