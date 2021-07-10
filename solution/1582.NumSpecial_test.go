package solution

import "testing"

func TestNumSpecial(t *testing.T) {
	mat := [][]int{
		{1, 0, 0},
		{0, 0, 1},
		{1, 0, 0},
	}
	exp := 1
	if NumSpecial(mat) != exp {
		t.Fail()
	}
}

func TestNumSpecial2(t *testing.T) {
	mat := [][]int{
		{0, 0, 0, 1},
		{1, 0, 0, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	}
	exp := 2
	if NumSpecial(mat) != exp {
		t.Fail()
	}
}

func TestNumSpecial3(t *testing.T) {
	mat := [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	exp := 3
	if NumSpecial(mat) != exp {
		t.Fail()
	}
}
