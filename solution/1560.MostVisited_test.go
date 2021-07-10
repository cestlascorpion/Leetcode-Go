package solution

import "testing"

func TestMostVisited(t *testing.T) {
	n := 4
	rounds := []int{1, 3, 1, 2}
	exp := []int{1, 2}
	res := MostVisited(n, rounds)
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}

func TestMostVisited2(t *testing.T) {
	n := 2
	rounds := []int{2,1,2,1,2,1,2,1,2}
	exp := []int{2}
	res := MostVisited(n, rounds)
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}

func TestMostVisited3(t *testing.T) {
	n := 7
	rounds := []int{1,3,5,7}
	exp := []int{1,2,3,4,5,6,7}
	res := MostVisited(n, rounds)
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}
