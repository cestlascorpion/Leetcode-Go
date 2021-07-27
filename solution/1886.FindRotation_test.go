package solution

import "testing"

func TestFindRotation(t *testing.T) {
	mat := [][]int{{0, 1}, {1, 0}}
	target := [][]int{{1, 0}, {0, 1}}
	if !FindRotation(mat, target) {
		t.Fail()
	}
}
