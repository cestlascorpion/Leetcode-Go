package solution

import "testing"

func TestNearestValidPoint(t *testing.T) {
	x := 3
	y := 4
	points := [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}}
	exp := 2
	if NearestValidPoint(x, y, points) != exp {
		t.Fail()
	}
}

func TestNearestValidPoint2(t *testing.T) {
	x := 3
	y := 4
	points := [][]int{{3, 4}}
	exp := 0
	if NearestValidPoint(x, y, points) != exp {
		t.Fail()
	}
}

func TestNearestValidPoint3(t *testing.T) {
	x := 3
	y := 4
	points := [][]int{{2, 3}}
	exp := -1
	if NearestValidPoint(x, y, points) != exp {
		t.Fail()
	}
}
