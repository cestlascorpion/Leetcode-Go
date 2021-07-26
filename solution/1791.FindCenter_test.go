package solution

import "testing"

func TestFindCenter(t *testing.T) {
	edges := [][]int{{1, 2}, {2, 3}, {4, 2}}
	exp := 2
	if FindCenter(edges) != exp {
		t.Fail()
	}
}
