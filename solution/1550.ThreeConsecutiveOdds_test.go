package solution

import "testing"

func TestThreeConsecutiveOdds(t *testing.T) {
	arr := []int{2, 4, 61}
	if ThreeConsecutiveOdds(arr) {
		t.Fail()
	}
}
