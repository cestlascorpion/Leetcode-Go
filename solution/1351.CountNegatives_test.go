package solution

import "testing"

func TestCountNegatives(t *testing.T) {
	grid := [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}
	exp := 8
	res := CountNegatives(grid)
	if exp != res {
		t.Fail()
	}
}
