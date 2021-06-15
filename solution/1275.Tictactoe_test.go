package solution

import "testing"

func TestTictactoe(t *testing.T) {
	moves := [][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}}
	exp := "A"
	res := Tictactoe(moves)
	if res != exp {
		t.Fail()
	}
}
