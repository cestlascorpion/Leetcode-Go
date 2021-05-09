package solution

import "testing"

func TestCountGoodRectangles(t *testing.T) {
	rectangles := [][]int{{2, 3}, {3, 7}, {4, 3}, {3, 7}}
	exp := 3
	if CountGoodRectangles(rectangles) != exp {
		t.Fail()
	}
}
