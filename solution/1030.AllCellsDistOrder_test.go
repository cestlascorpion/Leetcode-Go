package solution

import "testing"

func TestAllCellsDistOrder(t *testing.T) {
	R := 1
	C := 2
	r0 := 0
	c0 := 0
	exp := [][]int{{0, 0}, {0, 1}}
	res := AllCellsDistOrder(R, C, r0, c0)
	if len(exp) != len(res) {
		t.Fail()
	}
	for i := range res {
		if res[i][0] != exp[i][0] || res[i][1] != exp[i][1] {
			t.Fail()
		}
	}
}
