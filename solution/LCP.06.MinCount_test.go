package solution

import "testing"

func TestMinCount(t *testing.T) {
	coins := []int{4, 2, 1}
	if 4 != MinCount(coins) {
		t.Fail()
	}
}
