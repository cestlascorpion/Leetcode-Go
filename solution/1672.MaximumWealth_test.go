package solution

import "testing"

func TestMaximumWealth(t *testing.T) {
	a1 := [][]int{{1, 2, 3}, {3, 2, 1}}
	if ret := maximumWealth(a1); ret != 6 {
		t.Fatalf("exp 6, ret %d", ret)
	}

	a2 := [][]int{{1, 5}, {7, 3}, {3, 5}}
	if ret := maximumWealth(a2); ret != 10 {
		t.Fatalf("exp 6, ret %d", ret)
	}
}
