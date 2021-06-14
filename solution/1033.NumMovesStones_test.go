package solution

import (
	"testing"
)

func TestNumMovesStones(t *testing.T) {
	a, b, c := 1, 2, 5
	exp := []int{1, 2}
	res := NumMovesStones(a, b, c)
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}

func TestNumMovesStones2(t *testing.T) {
	a, b, c := 4, 3, 2
	exp := []int{0, 0}
	res := NumMovesStones(a, b, c)
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}

func TestNumMovesStones3(t *testing.T) {
	a, b, c := 3, 5, 1
	exp := []int{1, 2}
	res := NumMovesStones(a, b, c)
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}
