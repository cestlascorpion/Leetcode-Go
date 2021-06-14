package solution

import "testing"

func TestIsBoomerang(t *testing.T) {
	points := [][]int{{1, 1}, {2, 3}, {3, 2}}
	if !IsBoomerang(points) {
		t.Fail()
	}
}

func TestIsBoomerang2(t *testing.T) {
	points := [][]int{{1, 1}, {2, 2}, {3, 3}}
	if IsBoomerang(points) {
		t.Fail()
	}
}
