package solution

import "testing"

func TestValidPath(t *testing.T) {
	n := 6
	edges := [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}
	start := 0
	end := 5
	if ValidPath(n, edges, start, end) {
		t.Fail()
	}
}
