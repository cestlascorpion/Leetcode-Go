package solution

import (
	"testing"
)

func TestShiftGrid(t *testing.T) {
	k := 6
	grid := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	exp := [][]int{{4, 5, 6}, {7, 8, 9}, {1, 2, 3}}
	res := ShiftGrid(grid, k)
	m, n := len(exp), len(exp[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if res[i][j] != exp[i][j] {
				t.Fail()
			}
		}
	}
}
