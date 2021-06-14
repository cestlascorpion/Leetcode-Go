package solution

import "testing"

func TestFindJudge(t *testing.T) {
	N := 4
	trust := [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}
	exp := 3
	if FindJudge(N, trust) != exp {
		t.Fail()
	}
}

func TestFindJudge2(t *testing.T) {
	N := 3
	trust := [][]int{{1, 3}, {2, 3}, {3, 1}}
	exp := -1
	if FindJudge(N, trust) != exp {
		t.Fail()
	}
}

func TestFindJudge3(t *testing.T) {
	N := 2
	trust := [][]int{{1, 2}}
	exp := 2
	if FindJudge(N, trust) != exp {
		t.Fail()
	}
}
