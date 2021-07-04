package solution

import "testing"

func TestCanMakeArithmeticProgression(t *testing.T) {
	arr := []int{3, 5, 1}
	if !CanMakeArithmeticProgression(arr) {
		t.Fail()
	}
}
